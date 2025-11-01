package handlers

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"topinambur02.com/m/v2/senders"
	"topinambur02.com/m/v2/utils"
)

func HandleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	switch message.Text {
	case "/start":
		senders.SendWelcomeMessage(bot, message.Chat.ID)
	case "Показать текущее расписание":
		lessons, err := utils.LoadSchedule()
		if err != nil {
			panic("Error loading schedule!")
		}
		senders.SendSchedule(bot, message.Chat.ID, lessons)
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "Error: Unknown command")
		msg.ReplyToMessageID = message.MessageID
		if _, err := bot.Send(msg); err != nil {
			log.Printf("Failed to send schedule message: %v", err)
		}
	}
}
