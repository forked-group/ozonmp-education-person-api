package person

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
)

func (c commander) Help(inputMsg *tgbotapi.Message) {
	const op = "commander.Help"

	commandSuffix := "__" + c.domain + "__" + c.subdomain
	outputText := strings.Builder{}

	// TODO: команда сома должна рассказывать о себе
	outputText.WriteString("/help" + commandSuffix + " — print list of commands\n")
	outputText.WriteString("/list" + commandSuffix + " — get a list of entity\n")
	outputText.WriteString("/get" + commandSuffix + " <id> — get a entity with id\n")
	outputText.WriteString("/delete" + commandSuffix + " <id> — delete an existing entity\n")
	outputText.WriteString("/new" + commandSuffix + " <name> — create a new entity\n")
	outputText.WriteString("/edit" + commandSuffix + " <id> <name> — edit a entity\n")

	outputMsg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputText.String())

	if _, err := c.bot.Send(outputMsg); err != nil {
		log.Printf("%s: can't send message to chat: %v", op, err)
	}
}
