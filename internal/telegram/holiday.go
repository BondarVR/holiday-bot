package telegram

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"net/http"
	"strings"
	"time"
)

type Message struct {
	Name        string `json:"name"`
	NameLocal   string `json:"name_local"`
	Language    string `json:"language"`
	Description string `json:"description"`
	Country     string `json:"country"`
	Location    string `json:"location"`
	Type        string `json:"type"`
	Date        string `json:"date"`
	DateYear    string `json:"date_year"`
	DateMonth   string `json:"date_month"`
	DateDay     string `json:"date_day"`
	WeekDay     string `json:"week_day"`
}

func (b *Bot) getHolidayInfo(country string, message *tgbotapi.Message) (tgbotapi.MessageConfig, error) {
	timeT := b.parseDate(message.Date)
	requestURL := fmt.Sprintf("https://holidays.abstractapi.com/v1/?api_key=%s&country=%s&year=%d&month=%d&day=%d",
		b.cfg.ApiHoliday, country, timeT.Year(), timeT.Month(), timeT.Day())

	resp, err := http.Get(requestURL)
	if err != nil {
		b.lgr.Fatal(err)
	}
	defer resp.Body.Close()

	value, err := b.unmarshalJSON(resp)
	if err != nil {
		b.lgr.Fatal(err)
	}

	info := b.handleResponseForClient(value)
	msg := tgbotapi.NewMessage(message.Chat.ID, info)
	return msg, nil
}

func (b *Bot) handleResponseForClient(tmp []Message) string {
	var x []string
	for _, value := range tmp {
		x = append(x, value.Name)
	}
	finalString := strings.Join(x, ",\n")
	if finalString == "" {
		finalString = "There is no holiday in this country today :("
	}
	return finalString
}

func (b *Bot) unmarshalJSON(response *http.Response) ([]Message, error) {
	var v []Message
	body, err := io.ReadAll(response.Body)
	if err != nil {
		b.lgr.Fatal(err)
	}
	if err := json.Unmarshal(body, &v); err != nil {
		b.lgr.Fatal(err)
	}
	return v, nil
}

func (b *Bot) parseDate(date int) time.Time {
	dateConvert := int64(date)
	timeT := time.Unix(dateConvert, 0)
	return timeT
}
