package person

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c commander) Get(inputMsg *tgbotapi.Message) {
	const op = "commander.Get"

	argsStr := inputMsg.CommandArguments()
	args, err := parseArguments(argsStr)

	if err != nil {
		log.Printf("%s: can't parse arguments %q: %v", op, argsStr, err)
		return
	}

	if len(args) == 0 {
		log.Printf("%s: argument required", op)
		return
	}

	if len(args) > 1 {
		log.Printf("%s: too many arguments, want one: %q", op, argsStr)
		return
	}

	id, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		log.Printf("%s: invalid argument '%s': %v", op, args[0], err)
		return
	}

	item, err := c.service.Describe(id)
	if err != nil {
		log.Printf("%s: can't get item with id %d: %v", op, id, err)
		return
	}

	outputMsg := tgbotapi.NewMessage(inputMsg.Chat.ID, item.String())

	if _, err := c.bot.Send(outputMsg); err != nil {
		log.Printf("%s: can't send message to chat: %v", op, err)
	}
}
