package senders

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"topinambur02.com/m/v2/model"
	"topinambur02.com/m/v2/utils"
)

func SendSchedule(bot *tgbotapi.BotAPI, chatID int64, lessons []model.Lesson) {
	if len(lessons) == 0 {
		msg := tgbotapi.NewMessage(chatID, "На сегодня занятий нет 🎉")

		if _, err := bot.Send(msg); err != nil {
			log.Printf("Failed to send schedule message: %v", err)
		}

		return
	}

	var messageParts []string
	for i, lesson := range lessons {
		lessonText := utils.FormatLesson(lesson)
		messageParts = append(messageParts, fmt.Sprintf("%d. %s", i+1, lessonText))

		if (i+1)%5 == 0 || i == len(lessons)-1 {
			msg := tgbotapi.NewMessage(chatID, strings.Join(messageParts, "\n\n"))
			msg.ParseMode = "HTML"
			if _, err := bot.Send(msg); err != nil {
				log.Printf("Failed to send schedule message: %v", err)
			}
			messageParts = []string{}
		}
	}
}
