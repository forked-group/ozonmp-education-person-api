package person

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service"

	"log"
	"strings"
)

type callbackListData struct {
	Cursor uint64 `json:"cursor"`
}

const listPageSize = 5

func (c commander) List(inputMsg *tgbotapi.Message) {
	c.listNext(inputMsg.Chat.ID, 0)
}

func (c commander) ListCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	const op = "commander.ListCallback"

	var data callbackListData
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &data)
	if err != nil {
		log.Printf("%s: can't unmarshal calback data: %v", op, err)
		return
	}

	c.listNext(callback.Message.Chat.ID, data.Cursor)
}

func (c commander) listNext(chatID int64, cursor uint64) {
	const op = "commander.listNext"

	entries, err := c.service.List(cursor, listPageSize)

	switch {
	case err == service.ErrNotFound || len(entries) == 0:
		c.SendError(chatID, "no more entries")
		return
	case err != nil:
		log.Printf("%s: can't list entries: %v", op, err)
		c.SendError(chatID, "internal error")
		return
	}

	outputText := strings.Builder{}

	for _, person := range entries {
		outputText.WriteString(person.String())
		outputText.WriteByte('\n')
	}

	outputMsg := tgbotapi.NewMessage(chatID, outputText.String())

	serializedData, _ := json.Marshal(callbackListData{
		Cursor: entries[len(entries)-1].ID + 1,
	})

	var callbackPath = path.CallbackPath{
		Domain:       c.domain,
		Subdomain:    c.subdomain,
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	outputMsg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	if _, err = c.bot.Send(outputMsg); err != nil {
		log.Printf("%s: can't send message to chat: %v", op, err)
	}
}
