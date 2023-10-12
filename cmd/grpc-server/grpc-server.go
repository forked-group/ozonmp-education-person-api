package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/app/retranslator"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/app/sender"
	"time"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	repo2 "github.com/aaa2ppp/ozonmp-education-person-api/internal/app/repo"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/config"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/database"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/repo"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/server"
)

const (
	envDummyRepo = "DUMMY_REPO"
	batchSize    = 2 // TODO: ???
	kafka        = "localhost:9094"
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

	db := openDB(cfg.Database)
	if db == nil {
		log.Error().Msg("Can't open data base")
		return
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Error().Err(err).Msg("Can't close database")
			return
		}
		log.Info().Msg("Data base closed")
	}()

	eventRepo := repo2.EventRepoAdapter{Repo: repo.NewEventRepo(db)}
	eventSender, err := sender.NewEventSender(kafka)
	if err != nil {
		log.Error().Err(err).Msg("Can't create event sender")
		return
	}
	defer func() {
		err := eventSender.Close()
		if err != nil {
			log.Error().Err(err).Msg("Can't close event sender")
			return
		}
		log.Info().Msg("Event sender closed")
	}()

	eventRetranslator := startRetranslator(eventRepo, eventSender)
	defer func() {
		eventRetranslator.Close()
		log.Info().Msg("Event retranslator stopped")
	}()
	log.Info().Msg("Event retranslator started")

	personRepo := repo.NewPersonRepo(db, batchSize)
	if err := server.NewGrpcServer(personRepo).Start(&cfg); err != nil {
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

func startRetranslator(repo repo2.EventRepo, sender sender.EventSender) *retranslator.Retranslator {
	cfg := retranslator.Config{
		ChannelSize: 0,

		ConsumerCount:  1,
		ConsumeSize:    10,
		ConsumeTimeout: 1000 * time.Millisecond,

		ProducerCount:  10,
		ProduceTimeout: 1000 * time.Millisecond,

		CollectSize:     10,
		CollectMaxDelay: 1000 * time.Millisecond,

		WorkerCount:      2,
		WorkErrorTimeout: 1000 * time.Millisecond,

		Repo:   repo,
		Sender: sender,
	}

	return cfg.Start(context.TODO())
}
