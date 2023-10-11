package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/config"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/database"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/repo"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"strconv"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	_ = godotenv.Load() // XXX

	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	db := openDB(cfg.Database)
	if db == nil {
		log.Error().Msg("Can't open data base")

		return
	}
	defer db.Close()

	r := repo.NewEventRepo(db)

	if len(os.Args) < 2 {
		usage()
	}

	ctx := context.TODO()

	switch cmd, args := os.Args[1], os.Args[2:]; cmd {
	default:
		usage()

	case "lock":
		if len(args) != 1 {
			die("%s: number of records required", cmd)
		}

		n, err := strconv.Atoi(args[0])
		if err != nil {
			die("%s: can't parse number of records: %v", cmd, err)
		}

		events, err := r.Lock(ctx, uint64(n))
		if err != nil {
			die("%s: can't lock: %v", cmd, err)
		}

		info("%s: %d events locked", cmd, len(events))

		for _, e := range events {
			out("%+v\n", e)
		}

	case "unlock":
		if len(args) == 0 {
			warn("%s: list of event IDs required", cmd)
		}

		eventIDs := make([]uint64, len(args))

		for i, a := range args {
			id, err := strconv.Atoi(a)
			if err != nil {
				die("%s: can't parse id%d: %v", cmd, i+1, err)
			}
			eventIDs[i] = uint64(id)
		}

		n, err := r.Unlock(ctx, eventIDs)
		if err != nil {
			die("%s: can't unlock events: %v", cmd, err)
		}

		info("%s: %d events unlocked", cmd, n)

	case "remove":
		if len(args) == 0 {
			warn("%s: list of event IDs required", cmd)
		}

		eventIDs := make([]uint64, len(args))

		for i, a := range args {
			id, err := strconv.Atoi(a)
			if err != nil {
				die("%s: can't parse id%d: %v", cmd, i+1, err)
			}
			eventIDs[i] = uint64(id)
		}

		n, err := r.Remove(ctx, eventIDs)
		if err != nil {
			die("%s: can't remove events: %v", cmd, err)
		}

		info("%s: %d events removed", cmd, n)
	}
}

func out(f string, a ...any) {
	fmt.Printf(f, a...)
}

func msg(f string, a ...any) {
	if len(f) > 0 && f[len(f)-1] != '\n' {
		f += "\n"
	}
	fmt.Fprintf(os.Stderr, f, a...)
}

func debug(f string, a ...any) {
	msg(f, a...)
}

func info(f string, a ...any) {
	msg(f, a...)
}

func warn(f string, a ...any) {
	msg(f, a...)
}

func die(f string, a ...any) {
	msg(f, a...)
	os.Exit(1)
}

func usage() {
	die("Usage: %s {lock|unlock|remove} args...\n", filepath.Base(os.Args[0]))
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
