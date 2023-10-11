package person

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/rs/zerolog/log"
	"strings"
)

const listUsage = "/list%s"
const listPageSize = 5

func (c Commander) List(inputMsg *tgbotapi.Message) {
	c.listNext(inputMsg.Chat.ID, 0)
}

type listCallbackData struct {
	Cursor uint64 `json:"cursor"`
}

func (c Commander) listCallback(callback *tgbotapi.CallbackQuery, callbackPath callbackPath) {
	const op = "Commander.listCallback"

	var data listCallbackData
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &data)
	if err != nil {
		log.Printf("%s: can't unmarshal callback data: %v", op, err)
		return
	}

	c.listNext(callback.Message.Chat.ID, data.Cursor)
}

func (c Commander) listNext(chatID int64, cursor uint64) {
	const op = "Commander.listNext"

	entries, err := c.service.List(cursor, listPageSize)

	if err != nil {
		log.Printf("%s: can't list entries: %v", op, err)
		c.sendError(chatID, "internal error")
		return
	}

	if len(entries) == 0 {
		c.sendError(chatID, "no more entries")
		return
	}

	outputText := strings.Builder{}

	for _, person := range entries {
		outputText.WriteString(person.String())
		outputText.WriteByte('\n')
	}

	outputMsg := tgbotapi.NewMessage(chatID, outputText.String())

	serializedData, _ := json.Marshal(listCallbackData{
		Cursor: entries[len(entries)-1].ID + 1,
	})

	var path = callbackPath{
		Domain:       c.domain,
		Subdomain:    c.subdomain,
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	outputMsg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", path.String()),
		),
	)

	if _, err = c.bot.Send(outputMsg); err != nil {
		log.Printf("%s: can't send message to chat: %v", op, err)
	}
}
