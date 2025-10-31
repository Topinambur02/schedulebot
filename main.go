package main

import (
	"log"
	"time"

	"topinambur02.com/m/v2/configs"
	"topinambur02.com/m/v2/db"
	"topinambur02.com/m/v2/handlers"
	"topinambur02.com/m/v2/model"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	config := configs.NewConfig()
	db := db.InitDB()
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
			tgId := update.Message.From.ID
			username := update.Message.From.UserName
			messageText := update.Message.Text
			user := model.User{Username: username, Tg_id: tgId, CreatedAt: time.Now()}

			db.Create(&user)
			log.Printf("User: %s, Message: %s", username, messageText)
			
			go handlers.HandleMessage(bot, update.Message)
		}
	}
}