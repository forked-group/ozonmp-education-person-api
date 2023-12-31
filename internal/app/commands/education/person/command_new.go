package person

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rs/zerolog/log"
)

const newUsage = "/new%s [[first_name [middle_name]] last_name] [field=value ...]"

func (c Commander) New(inputMsg *tgbotapi.Message) {
	const op = "Commander.New"

	chatID := inputMsg.Chat.ID

	args, err := splitIntoArguments(inputMsg.CommandArguments())
	if err != nil {
		c.sendError(chatID, err.Error())
		return
	}

	if len(args) == 0 {
		c.sendError(chatID, "you must specify the field(s) of person")
		return
	}

	var (
		p person
		f personField // stub
	)

	if args, err = parsePersonNames(args, &p); err != nil {
		c.sendError(chatID, err.Error())
		return
	}

	if err = parsePersonFields(args, &p, &f); err != nil {
		c.sendError(chatID, err.Error())
		return
	}

	id, err := c.service.Create(p)

	if err != nil {
		log.Printf("%s: can't create person: %v", op, err)
		c.sendError(chatID, "internal error")
		return
	}

	c.sendOk(chatID, "new person %d created", id)
}
