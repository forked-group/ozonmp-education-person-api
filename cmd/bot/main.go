package main

import (
	personCommands "github.com/aaa2ppp/ozonmp-education-person-api/internal/app/commands/education/person"
	personService "github.com/aaa2ppp/ozonmp-education-person-api/internal/service/education/person"
	"log"
	"os"

	routerPkg "github.com/aaa2ppp/ozonmp-education-person-api/internal/app/router"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	token, found := os.LookupEnv("TOKEN")
	if !found {
		log.Panic("environment variable TOKEN not found in .env")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	// Uncomment if you want debugging
	// bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	router := routerPkg.NewRouter(bot)

	var service *personService.DummyPersonService
	if _, ok := os.LookupEnv("WITH_TEST_DATA"); ok {
		service = personService.NewDummyPersonServiceWithTestData()
	} else {
		service = personService.NewDummyPersonService()
	}

	router.SetRoute("education", "person",
		personCommands.NewPersonCommander(bot, service),
	)

	for update := range updates {
		router.HandleUpdate(update)
	}
}
