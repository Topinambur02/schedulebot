package main

import (
	"log"

	"topinambur02.com/m/v2/configs"
	"topinambur02.com/m/v2/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	config := configs.NewConfig()
	bot, err := tgbotapi.NewBotAPI(config.TELEGRAM_API_TOKEN)

	if err != nil {
		panic(err)
	}

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	log.Println("Bot started")

	for update := range updates {
		if update.Message != nil {
			username := update.Message.From.UserName
			messageText := update.Message.Text
			log.Printf("User: %s, Message: %s", username, messageText)
			go handlers.HandleMessage(bot, update.Message)
		}
	}
}