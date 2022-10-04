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
		msg, err := b.getHolidayInfo("ua", message)
		if err != nil {
			return err
		}
		if _, err := b.bot.Send(msg); err != nil {
			return err
		}
	case textEngland:
		msg, err := b.getHolidayInfo("gb", message)
		if err != nil {
			return err
		}
		if _, err := b.bot.Send(msg); err != nil {
			return err
		}
	case textFrance:
		msg, err := b.getHolidayInfo("fr", message)
		if err != nil {
			return err
		}
		if _, err := b.bot.Send(msg); err != nil {
			return err
		}
	case textPoland:
		msg, err := b.getHolidayInfo("pl", message)
		if err != nil {
			return err
		}
		if _, err := b.bot.Send(msg); err != nil {
			return err
		}
	case textUSA:
		msg, err := b.getHolidayInfo("us", message)
		if err != nil {
			return err
		}
		if _, err := b.bot.Send(msg); err != nil {
			return err
		}
	case textGerman:
		msg, err := b.getHolidayInfo("de", message)
		if err != nil {
			return err
		}
		if _, err := b.bot.Send(msg); err != nil {
			return err
		}
	default:
		return errors.New("Invalid type of text.")
	}
	return nil
}
