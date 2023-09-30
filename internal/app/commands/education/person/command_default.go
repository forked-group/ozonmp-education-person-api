package person

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c commander) Default(inputMsg *tgbotapi.Message) {
	const op = "commander.Default"

	log.Printf("[%s] %s", inputMsg.From.UserName, inputMsg.Text)
	c.Send(inputMsg.Chat.ID, "You wrote: "+inputMsg.Text)
}
