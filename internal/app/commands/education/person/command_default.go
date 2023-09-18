package person

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c commander) Default(inputMsg *tgbotapi.Message) {
	const op = "commander.Default"

	log.Printf("[%s] %s", inputMsg.From.UserName, inputMsg.Text)

	outputMsg := tgbotapi.NewMessage(inputMsg.Chat.ID, "You wrote: "+inputMsg.Text)

	if _, err := c.bot.Send(outputMsg); err != nil {
		log.Printf("%s: can't send message to chat: %v", op, err)
	}
}
