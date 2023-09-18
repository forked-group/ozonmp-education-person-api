package person

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func (c commander) Delete(inputMsg *tgbotapi.Message) {
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

	_, err = c.service.Remove(id)
	if err != nil {
		log.Printf("%s: can't delete item with id %d: %v", op, id, err)
		return
	}
}
