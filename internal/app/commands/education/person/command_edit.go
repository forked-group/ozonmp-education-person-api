package person

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service"
	"log"
	"strconv"
)

func (c commander) Edit(inputMsg *tgbotapi.Message) {
	const op = "commander.Edit"
	const usage = "/edit%s id field=value ..."

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

	id, err := strconv.ParseUint(args[0], 0, 64)
	if err != nil {
		c.SendError(chatID, "id must be a positive number")
		return
	}

	if len(args) == 1 {
		c.SendError(chatID, "you must specify the field(s) to update")
		return
	}

	p, err := c.service.Describe(id)
	switch {
	case err == service.ErrNotFound:
		c.SendError(chatID, "person id not found")
		return
	case err != nil:
		log.Printf("%s: can't get person: %v", op, err)
		c.SendError(chatID, "internal error")
		return
	}

	if err = parsePersonFields(args[1:], p); err != nil {
		c.SendError(chatID, err.Error())
		return
	}

	err = c.service.Update(id, *p)
	if err != nil {
		log.Printf("%s: can't update person: %v", op, err)
		c.SendError(chatID, "internal error")
		return
	}

	c.SendOk(chatID, "successful updated")
}
