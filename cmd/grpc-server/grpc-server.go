package main

import (
	"flag"
	"fmt"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/aaa2ppp/ozonmp-education-person-api/internal/config"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/database"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/interfaces"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/repo"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/server"
)

const (
	envDummyRepo = "DUMMY_REPO"
	batchSize    = 2 // TODO: ???
)

func main() {
	defer func() {
		log.Info().Msg("finished")
	}()
	log.Info().Msg("starting")

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

	//tracing, err := tracer.NewTracer(&cfg)
	//if err != nil {
	//	log.Error().Err(err).Msg("Failed init tracing")
	//
	//	return
	//}
	//defer tracing.Close()

	var r interfaces.PersonRepo
	db := openDB(cfg.Database)
	if db == nil {
		log.Error().Msg("Can't open data base")
		return
	}
	defer db.Close()

	r = repo.NewPersonRepo(db, batchSize)

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
