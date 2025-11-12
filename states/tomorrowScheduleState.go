package states

import (
	"topinambur02.com/m/v2/model"
	"topinambur02.com/m/v2/senders"
	"topinambur02.com/m/v2/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TomorrowScheduleState struct{}

func (s *TomorrowScheduleState) Handle(bot *tgbotapi.BotAPI, message *tgbotapi.Message, user model.User) {
	lessons, err := utils.LoadScheduleForTomorrow()
	if err != nil {
		panic("Error loading schedule!")
	}
	senders.SendSchedule(bot, message.Chat.ID, lessons)
}