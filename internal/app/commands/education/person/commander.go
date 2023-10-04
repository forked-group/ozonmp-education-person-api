package person

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rs/zerolog/log"
)

type Commander struct {
	domain    string
	subdomain string
	bot       *tgbotapi.BotAPI
	service   personService
}

func NewCommander(service personService) *Commander {
	return &Commander{
		service: service,
	}
}

func (c *Commander) Config(cfg commanderCfg) {
	c.bot = cfg.BotAPI
	c.domain = cfg.Domain
	c.subdomain = cfg.Subdomain
}

func (c Commander) checkPath(op string, domain, subdomain string) bool {

	if domain != c.domain || subdomain != c.subdomain {
		log.Printf("%s: unknown path - %s/%s", op, domain, subdomain)
		return false
	}

	return true
}

func (c Commander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath callbackPath) {
	const op = "Commander.HandleCallback"

	if !c.checkPath(op, callbackPath.Domain, callbackPath.Subdomain) {
		return
	}

	switch callbackPath.CallbackName {
	case "list":
		c.ListCallback(callback, callbackPath)
	default:
		log.Printf("%s: unknown callback name: %s", op, callbackPath.CallbackName)
	}
}

func (c Commander) HandleCommand(inputMsg *tgbotapi.Message, commandPath commandPath) {
	const op = "Commander.HandleCommand"

	if !c.checkPath(op, commandPath.Domain, commandPath.Subdomain) {
		return
	}

	switch commandPath.CommandName {
	case "help":
		c.Help(inputMsg)
	case "list":
		c.List(inputMsg)
	case "get":
		c.Get(inputMsg)
	case "delete":
		c.Delete(inputMsg)
	case "new":
		c.New(inputMsg)
	case "edit":
		c.Edit(inputMsg)
	default:
		c.Default(inputMsg)
	}
}

func (c Commander) sendError(chatID int64, msg string, a ...any) {
	c.send(chatID, "ERR "+msg, a...)
}

func (c Commander) sendOk(chatID int64, msg string, a ...any) {
	c.send(chatID, "OK "+msg, a...)
}

func (c Commander) send(chatID int64, msg string, a ...any) {
	const op = "Commander.send"

	output := tgbotapi.NewMessage(chatID, fmt.Sprintf(msg, a...))

	if _, err := c.bot.Send(output); err != nil {
		log.Printf("%s: can't send message to chat: %v", op, err)
	}
}

func (c Commander) cmdSuffix() string {
	return "__" + c.domain + "__" + c.subdomain
}
