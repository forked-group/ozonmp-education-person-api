package person

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service"
	"log"
	"strconv"
)

func (c commander) Get(inputMsg *tgbotapi.Message) {
	const op = "commander.Get"
	const usage = "usage: /get%s id"

	chatID := inputMsg.Chat.ID

	args, err := splitIntoArguments(inputMsg.CommandArguments())
	if err != nil {
		c.SendError(chatID, err.Error())
		return
	}

	if len(args) == 0 {
		c.SendError(chatID, "you must specify the person id")
		return
	}

	if len(args) > 1 {
		c.SendError(chatID, "you can only specify one person id")
		return
	}

	id, err := strconv.ParseUint(args[0], 0, 64)
	if err != nil {
		c.SendError(chatID, "id must be a positive number")
		return
	}

	person, err := c.service.Describe(id)

	switch {
	case err == service.ErrNotFound:
		c.SendError(chatID, "person id not found")
		return
	case err != nil:
		log.Printf("%s: can't get person: %v", op, err)
		c.SendError(chatID, "internal error")
		return
	}

	c.SendOk(chatID, person.String())
}
