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
	_ = db.InitClient()
	sqliteDB := db.InitDB()
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
			var user model.User
			result := sqliteDB.Db.Where("tg_id = ?", tgId).First(&user)
			
			if result.Error != nil {
				user = model.User{
					Username:  username,
					Tg_id:     tgId,
					CreatedAt: time.Now(),
					Role: "USER",
				}
				sqliteDB.Db.Create(&user)
				log.Printf("New user created: %s (ID: %d)", username, tgId)
			} else {
				if user.Username != username {
					user.Username = username
					sqliteDB.Db.Save(&user)
				}
			}

			log.Printf("User: %s, Message: %s", username, messageText)

			go handlers.HandleMessage(bot, update.Message, user)
		}
	}
}
