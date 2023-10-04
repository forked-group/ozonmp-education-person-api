package person

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)

func (c Commander) Help(inputMsg *tgbotapi.Message) {
	const op = "Commander.Help"

	suffix := c.cmdSuffix()
	msg := strings.Builder{}
	msg.Grow(512)

	// TODO: команда сома должна рассказывать о себе

	msg.WriteString("/help")
	msg.WriteString(suffix)
	msg.WriteString(" — print list of commands\n")

	msg.WriteString("/list")
	msg.WriteString(suffix)
	msg.WriteString(" — get a list of entity\n")

	msg.WriteString("/get")
	msg.WriteString(suffix)
	msg.WriteString(" id — get a entity with id\n")

	msg.WriteString("/delete")
	msg.WriteString(suffix)
	msg.WriteString(" id — delete an existing entity\n")

	msg.WriteString("/new")
	msg.WriteString(suffix)
	msg.WriteString(" [[first_name [middle_name]] last_name] [field=value ...] — create a new entity\n")

	msg.WriteString("/edit")
	msg.WriteString(suffix)
	msg.WriteString(" id [field=value ...] — edit a entity\n")

	c.send(inputMsg.Chat.ID, msg.String())
}
