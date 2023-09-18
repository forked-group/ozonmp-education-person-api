package person

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/education"
	"log"
	"strconv"
	"strings"
)

func (c commander) Edit(inputMsg *tgbotapi.Message) {
	const op = "commander.Edit"

	args := strings.SplitN(inputMsg.CommandArguments(), " ", 2)
	if len(args) != 2 {
		log.Printf("two arguments are required")
		return
	}

	id, err := strconv.ParseUint(args[0], 0, 64)
	if err != nil {
		log.Printf("%s: can't parse id '%s': %v", op, args[0], err)
		return
	}

	name := strings.TrimSpace(args[1])
	if len(name) == 0 {
		log.Printf("%s: argument required", op)
		return
	}

	err = c.service.Update(id, education.Person{Name: name})
	if err != nil {
		log.Printf("%s: can't update the item with id %d: %v", op, id, err)
		return
	}

	outputMsg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		fmt.Sprintf("Item updated: %d", id),
	)

	if _, err := c.bot.Send(outputMsg); err != nil {
		log.Printf("%s: can't send message to chat: %v", op, err)
	}
}
