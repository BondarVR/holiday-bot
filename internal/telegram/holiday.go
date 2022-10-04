package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"log"
	"net/http"
	"time"
)

type Message struct {
	Name string `json:"name"`
}

func (b *Bot) getHolidayInfo(country string, message *tgbotapi.Message) (tgbotapi.MessageConfig, error) {
	timeT := b.parseDate(message.Date)
	requestURL := fmt.Sprintf("https://holidays.abstractapi.com/v1/?api_key=%s&country=%s&year=%d&month=%d&day=%d",
		b.cfg.ApiHoliday, country, timeT.Year(), timeT.Month(), timeT.Day())

	resp, err := http.Get(requestURL)
	if err != nil {
		b.lgr.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, string(body))

	return msg, nil
}

func (b *Bot) parseDate(date int) time.Time {
	dateConvert := int64(date)
	timeT := time.Unix(dateConvert, 0)
	return timeT
}
