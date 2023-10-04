package person

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c Commander) Edit(inputMsg *tgbotapi.Message) {
	const op = "Commander.Edit"
	const usage = "/edit%s id field=value ..."

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

	id, err := strconv.ParseUint(args[0], 0, 64)
	if err != nil {
		c.sendError(chatID, "id must be a positive number")
		return
	}

	if len(args) == 1 {
		c.sendError(chatID, "you must specify the field(s) to update")
		return
	}

	p, err := c.service.Describe(id)
	if err != nil {
		log.Printf("%s: can't get person: %v", op, err)
		c.sendError(chatID, "internal error")
		return
	}

	if p == nil {
		c.sendError(chatID, "person id not found")
		return
	}

	if err = parsePersonFields(args[1:], p); err != nil {
		c.sendError(chatID, err.Error())
		return
	}

	ok, err := c.service.Update(id, *p)
	if err != nil {
		log.Printf("%s: can't update person: %v", op, err)
		c.sendError(chatID, "internal error")
		return
	}

	if !ok {
		log.Printf("%s: person %d not found: %v", op, id, err)
		c.sendError(chatID, "person not found")
		return
	}

	c.sendOk(chatID, "successful updated")
}
