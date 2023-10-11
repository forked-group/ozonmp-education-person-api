package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/aaa2ppp/ozonmp-education-person-api/internal/interfaces"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	personCommands "github.com/aaa2ppp/ozonmp-education-person-api/internal/app/commands/education/person"
	routerPkg "github.com/aaa2ppp/ozonmp-education-person-api/internal/app/router"
	personService "github.com/aaa2ppp/ozonmp-education-person-api/internal/service/education/person"
	pb "github.com/aaa2ppp/ozonmp-education-person-api/pkg/education-person-api"
)

const (
	botTokenEnv = "TOKEN"
	domain      = "education"
	subdomain   = "person"
	//grpcServer = domain + "-" + subdomain + "-api:8082"
	grpcServer = "localhost:8082"
)

func main() {
	_ = godotenv.Load() // XXX to set environment variables from .env

	conn, err := grpc.Dial(grpcServer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error().Err(err).Msgf("can't dial to %s", grpcServer)
		return
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Error().Err(err).Msgf("can't close connection to %s", grpcServer)
		}
		log.Info().Msgf("connection to %s closed", grpcServer)
	}()
	log.Info().Msgf("successful created connection to %s", grpcServer)

	client := pb.NewEducationPersonApiServiceClient(conn)
	service := personService.NewService(client)

	router, err := startTelegramBot(service)
	if err != nil {
		log.Error().Err(err).Msgf("can't start telegram bot")
		return
	}
	defer func() {
		router.Close()
		log.Info().Msg("telegram bot stopped")
	}()
	log.Info().Msg("telegram bot started")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	sig := <-c
	log.Info().Msgf("%v signal received", sig)
}

func startTelegramBot(service interfaces.PersonService) (*routerPkg.Router, error) {
	const op = "startTelegramBot"

	token, found := os.LookupEnv(botTokenEnv)
	if !found {
		return nil, fmt.Errorf("%s: environment variable %s is not set", op, botTokenEnv)
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	// Uncomment if you want debugging
	// bot.Debug = true

	log.Info().Str("op", op).Msgf("authorized on account %s", bot.Self.UserName)

	router := routerPkg.NewRouter(bot)
	router.Route(domain, subdomain, personCommands.NewCommander(service))

	err = router.Start(context.Background(), tgbotapi.UpdateConfig{
		Timeout: 60, // WTF
	})
	if err != nil {
		return nil, fmt.Errorf("%s: can't start router: %w", op, err)
	}

	return router, nil
}
