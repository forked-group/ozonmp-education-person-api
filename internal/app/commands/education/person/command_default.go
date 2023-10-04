package person

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rs/zerolog/log"
)

func (c Commander) Default(inputMsg *tgbotapi.Message) {
	const op = "Commander.Default"

	log.Printf("[%s] %s", inputMsg.From.UserName, inputMsg.Text)
	c.send(inputMsg.Chat.ID, "You wrote: %s", inputMsg.Text)
}
