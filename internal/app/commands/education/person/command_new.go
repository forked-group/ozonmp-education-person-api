package person

import (
	"fmt"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/model/education"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func (c commander) New(inputMsg *tgbotapi.Message) {
	const op = "commander.New"
	const usage = "/new%s [[first_name [middle_name]] last_name] [field=value ...]"

	chatID := inputMsg.Chat.ID

	args, err := splitIntoArguments(inputMsg.CommandArguments())
	if err != nil {
		c.SendError(chatID, err.Error())
		return
	}

	if len(args) == 0 {
		c.SendError(chatID, "you must specify the field(s) of person")
		return
	}

	var p education.Person

	if args, err = parsePersonNames(args, &p); err != nil {
		c.SendError(chatID, err.Error())
		return
	}

	if err = parsePersonFields(args, &p); err != nil {
		c.SendError(chatID, err.Error())
		return
	}

	id, err := c.service.Create(p)

	if err != nil {
		log.Printf("%s: can't create person: %v", op, err)
		c.SendError(chatID, "internal error")
		return
	}

	c.SendOk(chatID, fmt.Sprintf("create new person with id=%d", id))
}
