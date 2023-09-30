package person

import (
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/app/router"
	service "github.com/aaa2ppp/ozonmp-education-person-api/internal/service/education/person"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// TODO: Кто использует этот интерфейс?
//  Почему я должен возвращать его в NewPersonCommander?
//  Я вижу, что нужен только router.Commander.

type PersonCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)  // return error not implemented - where? how?
	Edit(inputMsg *tgbotapi.Message) // return error not implemented - where? how?

	router.Commander
}

const (
	myDomain    = "education"
	mySubdomain = "person"
)

func NewPersonCommander(bot *tgbotapi.BotAPI, service service.PersonService) PersonCommander {
	return newCommander(myDomain, mySubdomain, bot, service)
}
