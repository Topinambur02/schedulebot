package utils

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendBotMessage(text string, message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Failed to send schedule message: %v", err)
	}
}
