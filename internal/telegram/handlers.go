package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
)

const (
	commandStart = "start"
	textUkraine  = "Ukraine"
	textGerman   = "German"
	textUSA      = "USA"
	textEngland  = "England"
	textPoland   = "Poland"
	textFrance   = "France"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		if err := b.handleStartCommand(message); err != nil {
			return err
		}
		return nil
	default:
		return errors.New("Invalid type of command.")
	}
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	var country = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(textUkraine),
			tgbotapi.NewKeyboardButton(textGerman),
			tgbotapi.NewKeyboardButton(textEngland),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(textPoland),
			tgbotapi.NewKeyboardButton(textUSA),
			tgbotapi.NewKeyboardButton(textFrance),
		),
	)
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyMarkup = country
	_, err := b.bot.Send(msg)
	return err
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	switch message.Text {
	case textUkraine:
		country := "ua"
		if err := b.handleSendMessage(country, message); err != nil {
			return err
		}
	case textEngland:
		country := "gb"
		if err := b.handleSendMessage(country, message); err != nil {
			return err
		}
	case textFrance:
		country := "fr"
		if err := b.handleSendMessage(country, message); err != nil {
			return err
		}
	case textPoland:
		country := "pl"
		if err := b.handleSendMessage(country, message); err != nil {
			return err
		}
	case textUSA:
		country := "us"
		if err := b.handleSendMessage(country, message); err != nil {
			return err
		}
	case textGerman:
		country := "de"
		if err := b.handleSendMessage(country, message); err != nil {
			return err
		}
	default:
		return errors.New("Invalid type of text.")
	}
	return nil
}

func (b *Bot) handleSendMessage(country string, message *tgbotapi.Message) error {
	msg, err := b.getHolidayInfo(country, message)
	if err != nil {
		return err
	}
	if _, err := b.bot.Send(msg); err != nil {
		return err
	}
	return nil
}
