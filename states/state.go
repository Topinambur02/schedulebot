package states

import (
	"topinambur02.com/m/v2/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type State interface {
	Handle(bot *tgbotapi.BotAPI, message *tgbotapi.Message, user model.User)
}