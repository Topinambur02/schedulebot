package main

import (
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
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message != nil {
			handlers.HandleMessage(bot, update.Message)
		}
	}
}