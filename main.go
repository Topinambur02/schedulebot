package main

import (
	"log"

	"topinambur02.com/m/v2/configs"
	"topinambur02.com/m/v2/db"
	"topinambur02.com/m/v2/service"
	"topinambur02.com/m/v2/states"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	config := configs.NewConfig()
	db.InitRedisClient()
	sqliteDB := db.InitDB()
	bot, err := tgbotapi.NewBotAPI(config.TELEGRAM_API_TOKEN)

	if err != nil {
		panic(err)
	}

	stateHandler := states.NewStateContext(bot)
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	log.Println("Bot started")

	for update := range updates {
		if update.Message != nil {
			tgId := update.Message.From.ID
			username := update.Message.From.UserName
			messageText := update.Message.Text
			service := service.UserService{Db: sqliteDB}
			user, err := service.GetOrCreateUser(tgId, username)

			if err != nil {
				log.Fatalln("Error while creating user")
			}

			log.Printf("User: %s, Message: %s", username, messageText)

			go stateHandler.HandleMessage(bot, update.Message, user)
		}
	}
}
