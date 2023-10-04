package person

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rs/zerolog/log"
	"strconv"
)

func (c Commander) Delete(inputMsg *tgbotapi.Message) {
	const op = "Commander.Delete"
	const usage = "/delete%s id"

	chatID := inputMsg.Chat.ID

	args, err := splitIntoArguments(inputMsg.CommandArguments())
	if err != nil {
		c.sendError(chatID, err.Error())
		return
	}

	if len(args) == 0 {
		c.sendError(chatID, "you must specify the person id")
		return
	}

	if len(args) > 1 {
		c.sendError(chatID, "you can only specify one person id")
		return
	}

	id, err := strconv.ParseUint(args[0], 0, 64)
	if err != nil {
		c.sendError(chatID, "id must be a positive number")
		return
	}

	ok, err := c.service.Remove(id)
	if err != nil {
		log.Printf("%s: can't remove person %d: %v", op, id, err)
		c.sendError(chatID, "internal error")
		return
	}

	if !ok {
		c.sendError(chatID, "person %d not found", id)
		return
	}

	c.sendOk(chatID, "person %d successful deleted", id)
}
