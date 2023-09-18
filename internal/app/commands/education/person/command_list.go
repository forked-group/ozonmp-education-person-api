package person

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"

	"log"
	"strings"
)

type callbackListData struct {
	Offset uint64 `json:"offset"`
}

const listPageSize = 5

func (c commander) List(inputMsg *tgbotapi.Message) {
	const op = "commander.List"
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

	c.listNext(callback.Message.Chat.ID, data.Offset)
}

func (c commander) listNext(chatID int64, cursor uint64) {
	const op = "commander.listNext"

	items, err := c.service.List(cursor, listPageSize)

	if err != nil {
		log.Printf("%s: can't get list of items: %v", op, err)
		return
	}

	if len(items) == 0 {
		log.Printf("%s: no any items", op)
		return
	}

	outputText := strings.Builder{}

	for _, it := range items {
		outputText.WriteString(it.String())
		outputText.WriteByte('\n')
	}

	outputMsg := tgbotapi.NewMessage(chatID, outputText.String())

	serializedData, _ := json.Marshal(callbackListData{
		Offset: items[len(items)-1].ID,
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
