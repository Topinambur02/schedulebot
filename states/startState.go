package states

import (
	"topinambur02.com/m/v2/model"
	"topinambur02.com/m/v2/senders"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type StartState struct{}

func (s *StartState) Handle(bot *tgbotapi.BotAPI, message *tgbotapi.Message, user model.User) {
	senders.SendWelcomeMessage(bot, message.Chat.ID, user.Role)
}