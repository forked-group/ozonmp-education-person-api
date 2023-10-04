package main

import (
	"context"
	"errors"
	personCommands "github.com/aaa2ppp/ozonmp-education-person-api/internal/app/commands/education/person"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/interfaces"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/repo"
	personService "github.com/aaa2ppp/ozonmp-education-person-api/internal/service/education/person"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"

	routerPkg "github.com/aaa2ppp/ozonmp-education-person-api/internal/app/router"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

const batchSize = 2

func main() {
	_ = godotenv.Load() // XXX

	router, err := startBot(personService.NewService(repo.NewDummyRepo(batchSize)))
	if err != nil {
		log.Error().Err(err).Msgf("Can't start Telegram bot")

		return
	}
	defer func() {
		router.Close()
		log.Info().Msg("Telegram bot stopped")
	}()

	log.Info().Msg("Telegram bot started")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Block until a signal is received.
	<-c
}

func startBot(service interfaces.PersonService) (*routerPkg.Router, error) {
	token, found := os.LookupEnv("TOKEN")
	if !found {
		return nil, errors.New("environment variable TOKEN not found in .env")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	// Uncomment if you want debugging
	// bot.Debug = true

	log.Info().Msgf("Authorized on account %s", bot.Self.UserName)

	router := routerPkg.NewRouter(bot)

	router.Route("education", "person",
		personCommands.NewCommander(service),
	)

	err = router.Start(context.Background(), tgbotapi.UpdateConfig{
		Timeout: 60,
	})

	return router, nil
}
