package states

import (
	"context"
	"log"

	"topinambur02.com/m/v2/cacheUtils"
	"topinambur02.com/m/v2/db"
	"topinambur02.com/m/v2/model"
	"topinambur02.com/m/v2/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type DeleteCacheState struct{}

func (s *DeleteCacheState) Handle(bot *tgbotapi.BotAPI, message *tgbotapi.Message, user model.User) {
	redis := db.GetClient()
	ctx := context.Background()
	err := cacheUtils.DeleteCache(redis, ctx)

	if err != nil {
		log.Fatalln("Failed to delete cache from redis")
		utils.SendBotMessage("❌ Failed to delete cache from redis", message, bot)
		return
	}

	utils.SendBotMessage("✅ Success", message, bot)
}