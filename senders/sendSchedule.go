package senders

import (
	"fmt"
	"strings"

	"topinambur02.com/m/v2/model"
	"topinambur02.com/m/v2/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendSchedule(bot *tgbotapi.BotAPI, chatID int64, lessons []model.Lesson) {
	if len(lessons) == 0 {
		msg := tgbotapi.NewMessage(chatID, "На сегодня занятий нет 🎉")
		bot.Send(msg)
		return
	}

	var messageParts []string
	for i, lesson := range lessons {
		lessonText := utils.FormatLesson(lesson)
		messageParts = append(messageParts, fmt.Sprintf("%d. %s", i+1, lessonText))
		
		if (i+1)%5 == 0 || i == len(lessons)-1 {
			msg := tgbotapi.NewMessage(chatID, strings.Join(messageParts, "\n\n"))
			msg.ParseMode = "HTML"
			bot.Send(msg)
			messageParts = []string{}
		}
	}
}