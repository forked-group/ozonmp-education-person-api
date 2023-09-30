package demo

import (
	"log"

	"github.com/aaa2ppp/ozonmp-education-person-api/internal/app/commands/demo/subdomain"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/app/path"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type DemoCommander struct {
	bot                *tgbotapi.BotAPI
	subdomainCommander Commander
}

func NewDemoCommander(
	bot *tgbotapi.BotAPI,
) *DemoCommander {
	return &DemoCommander{
		bot: bot,
		// subdomainCommander
		subdomainCommander: subdomain.NewDemoSubdomainCommander(bot),
	}
}

func (c *DemoCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "subdomain":
		c.subdomainCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("DemoCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *DemoCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "subdomain":
		c.subdomainCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("DemoCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}