package main

import (
	personCommands "github.com/aaa2ppp/ozonmp-education-person-api/internal/app/commands/education/person"
	routerPkg "github.com/aaa2ppp/ozonmp-education-person-api/internal/app/router"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/config"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/repo"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/server"
	personService "github.com/aaa2ppp/ozonmp-education-person-api/internal/service/education/person"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

var (
	batchSize uint = 2
)

func main() {
	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	//migration := flag.Bool("migration", true, "Defines the migration start option")
	//flag.Parse()

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

	//dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
	//	cfg.Database.Host,
	//	cfg.Database.Port,
	//	cfg.Database.User,
	//	cfg.Database.Password,
	//	cfg.Database.Name,
	//	cfg.Database.SslMode,
	//)
	//
	//db, err := database.NewPostgres(dsn, cfg.Database.Driver)
	//if err != nil {
	//	log.Fatal().Err(err).Msg("Failed init postgres")
	//}
	//defer db.Close()
	//
	//*migration = false // todo: need to delete this line for homework-4
	//if *migration {
	//	if err = goose.Up(db.DB, cfg.Database.Migrations); err != nil {
	//		log.Error().Err(err).Msg("Migration failed")
	//
	//		return
	//	}
	//}
	//
	//tracing, err := tracer.NewTracer(&cfg)
	//if err != nil {
	//	log.Error().Err(err).Msg("Failed init tracing")
	//
	//	return
	//}
	//defer tracing.Close()
	//
	//r := repo.NewRepo(db, batchSize)

	var service *personService.DummyPersonService
	if _, ok := os.LookupEnv("WITH_TEST_DATA"); ok {
		service = personService.NewDummyPersonServiceWithTestData()
	} else {
		service = personService.NewDummyPersonService()
	}

	go startBot(service)

	r := repo.NewDummyRepo(service) // TODO: должно быть: API [-> Service] -> Repo -> DB

	if err := server.NewGrpcServer(r).Start(&cfg); err != nil {
		log.Error().Err(err).Msg("Failed creating gRPC server")

		return
	}
}

func startBot(service personService.PersonService) {
	_ = godotenv.Load() // TODO: remove this

	token, found := os.LookupEnv("TOKEN")
	if !found {
		log.Panic().Msg("environment variable TOKEN not found in .env")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic().Err(err)
	}

	// Uncomment if you want debugging
	// bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic().Err(err)
	}

	router := routerPkg.NewRouter(bot)

	router.SetRoute("education", "person",
		personCommands.NewPersonCommander(bot, service),
	)

	for update := range updates {
		router.HandleUpdate(update)
	}
}
