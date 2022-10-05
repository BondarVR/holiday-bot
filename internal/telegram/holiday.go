package telegram

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"net/http"
	"time"
)

type Message struct {
	Message []struct {
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
}

func (b *Bot) getHolidayInfo(country string, message *tgbotapi.Message) (tgbotapi.MessageConfig, error) {
	timeT := b.parseDate(message.Date)
	requestURL := fmt.Sprintf("https://holidays.abstractapi.com/v1/?api_key=%s&country=%s&year=%d&month=%d&day=%d",
		b.cfg.ApiHoliday, country, timeT.Year(), timeT.Month(), timeT.Day())
	fmt.Println(requestURL)

	resp, err := http.Get(requestURL)
	if err != nil {
		b.lgr.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		b.lgr.Fatal(err)
	}

	tmp, err := b.unmarshalJSON(body)
	if err != nil {
		b.lgr.Fatal(err)
	}

	b.lgr.Info(tmp)

	msg := tgbotapi.NewMessage(message.Chat.ID, "лика пидор")

	return msg, nil
}

func (b *Bot) unmarshalJSON(data []byte) (Message, error) {
	var v Message
	if err := json.Unmarshal(data, &v); err != nil {
		b.lgr.Fatal(err)
	}
	return v, nil
}

func (b *Bot) parseDate(date int) time.Time {
	dateConvert := int64(date)
	timeT := time.Unix(dateConvert, 0)
	return timeT
}
