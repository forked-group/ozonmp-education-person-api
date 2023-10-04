package main

import (
	"flag"
	"fmt"
	personCommands "github.com/aaa2ppp/ozonmp-education-person-api/internal/app/commands/education/person"
	routerPkg "github.com/aaa2ppp/ozonmp-education-person-api/internal/app/router"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/config"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/database"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/interfaces"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/repo"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/server"
	personService "github.com/aaa2ppp/ozonmp-education-person-api/internal/service/education/person"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"sync"
)

var (
	batchSize uint = 2 // TODO: ???
)

func main() {
	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	log.Info().
		Str("version", cfg.Project.Version).
		Str("commitHash", cfg.Project.CommitHash).
		Bool("debug", cfg.Project.Debug).
		Str("environment", cfg.Project.Environment).
		Msgf("Starting service: %s", cfg.Project.Name)

	if cfg.Project.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	db := openDB(cfg.Database)
	if db == nil {
		log.Error().Msg("Can't open data base")
		return
	}
	defer db.Close()

	//tracing, err := tracer.NewTracer(&cfg)
	//if err != nil {
	//	log.Error().Err(err).Msg("Failed init tracing")
	//
	//	return
	//}
	//defer tracing.Close()

	//r := repo.NewDummyRepo(batchSize)
	r := repo.NewRepo(db, batchSize)

	stopBot := startBot(personService.NewService(r))
	defer stopBot()

	if err := server.NewGrpcServer(r).Start(&cfg); err != nil {
		log.Error().Err(err).Msg("Failed creating gRPC server")

		return
	}
}

func openDB(cfg config.Database) *sqlx.DB {
	migration := flag.Bool("migration", true, "Defines the migration start option")
	flag.Parse()

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name,
		cfg.SslMode,
	)

	db, err := database.NewPostgres(dsn, cfg.Driver)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed init postgres")
		return nil
	}

	*migration = false // todo: need to delete this line for homework-4
	if *migration {
		if err = goose.Up(db.DB, cfg.Migrations); err != nil {
			log.Error().Err(err).Msg("Migration failed")

			_ = db.Close()
			return nil
		}
	}

	return db
}

func startBot(service interfaces.PersonService) func() {
	_ = godotenv.Load() // XXX

	token, found := os.LookupEnv("TOKEN")
	if !found {
		log.Panic().Msg("environment variable TOKEN not found in .env") // TODO: return error
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic().Err(err) // TODO: return error
	}

	// Uncomment if you want debugging
	// bot.Debug = true

	log.Info().Msgf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic().Err(err) // TODO: return error
	}

	router := routerPkg.NewRouter(bot)

	router.Route("education", "person",
		personCommands.NewCommander(bot, service),
	)

	cancel := make(chan struct{})
	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			select {
			case update := <-updates:
				router.HandleUpdate(update)
			case <-cancel:
				return
			}
		}
	}()

	var once sync.Once

	return func() {
		once.Do(func() {
			close(cancel)
			<-done
		})
	}
}
