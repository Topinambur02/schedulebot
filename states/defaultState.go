package states

import (
	"topinambur02.com/m/v2/model"
	"topinambur02.com/m/v2/utils"
	
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type DefaultState struct{}

func (s *DefaultState) Handle(bot *tgbotapi.BotAPI, message *tgbotapi.Message, user model.User) {
	utils.SendBotMessage("Error: Unknown command", message, bot)
}