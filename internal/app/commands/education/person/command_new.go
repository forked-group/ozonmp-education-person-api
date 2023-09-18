package person

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/education"
	"log"
	"strings"
)

func (c commander) New(inputMsg *tgbotapi.Message) {
	const op = "commander.New"

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

	items := make([]education.Person, 0, len(args))

	for _, name := range args {
		it := education.Person{Name: name}
		id, err := c.service.Create(it)
		if err != nil {
			log.Printf("%s: can't create new item with name '%s': %v", op, name, err)
			// TODO: rollback?
			break
		}
		it.ID = id
		items = append(items, it)
	}

	outputText := strings.Builder{}

	for _, it := range items {
		outputText.WriteString(it.String())
		outputText.WriteByte('\n')
	}

	outputMsg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputText.String())

	if _, err := c.bot.Send(outputMsg); err != nil {
		log.Printf("%s: can't send message to chat: %v", op, err)
	}
}
