package person

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)

const helpUsage = "/help%s"

func (c Commander) Help(inputMsg *tgbotapi.Message) {
	suffix := c.cmdSuffix()

	msg := strings.Builder{}
	msg.Grow(512)

	// TODO: команда сома должна рассказывать все о себе

	msg.WriteString(fmt.Sprintf(helpUsage, suffix))
	msg.WriteString(" — print list of commands \n")

	msg.WriteString(fmt.Sprintf(getUsage, suffix))
	msg.WriteString(" — get a entity with id\n")

	msg.WriteString(fmt.Sprintf(listUsage, suffix))
	msg.WriteString(" — get a list of entity\n")

	msg.WriteString(fmt.Sprintf(deleteUsage, suffix))
	msg.WriteString(" — delete an existing entity\n")

	msg.WriteString(fmt.Sprintf(newUsage, suffix))
	msg.WriteString(" — create a new entity\n")

	msg.WriteString(fmt.Sprintf(editUsage, suffix))
	msg.WriteString(" — edit a entity\n")

	c.send(inputMsg.Chat.ID, msg.String())
}
