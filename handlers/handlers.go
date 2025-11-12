package handlers

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"topinambur02.com/m/v2/cacheUtils"
	"topinambur02.com/m/v2/db"
	"topinambur02.com/m/v2/model"
	"topinambur02.com/m/v2/senders"
	"topinambur02.com/m/v2/utils"
)

func HandleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message, user model.User) {
	switch message.Text {
	case "/start":
		senders.SendWelcomeMessage(bot, message.Chat.ID, user.Role)
	case "Показать текущее расписание":
		lessons, err := utils.LoadSchedule()
		if err != nil {
			panic("Error loading schedule!")
		}
		senders.SendSchedule(bot, message.Chat.ID, lessons)
	case "Удалить кэш для Redis":
		redis := db.GetClient()
		ctx := context.Background()
		err := cacheUtils.DeleteCache(redis, ctx)

		if err != nil {
			log.Fatalln("Failed to delete cache from redis")
			utils.SendBotMessage("❌ Failed to delete cache from redis", message, bot)
		}

		utils.SendBotMessage("✅ Success", message, bot)
	case "Посмотреть расписание на завтра":
		lessons, err := utils.LoadScheduleForTomorrow()
		if err != nil {
			panic("Error loading schedule!")
		}
		senders.SendSchedule(bot, message.Chat.ID, lessons)
	default:
		utils.SendBotMessage("Error: Unknown command", message, bot)
	}
}
