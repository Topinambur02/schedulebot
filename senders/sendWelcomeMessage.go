package senders

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendWelcomeMessage(bot *tgbotapi.BotAPI, chatID int64, role string) {
	text := "Добро пожаловать! Я бот с расписанием занятий.\nВыберите действие:"

    rows := [][]tgbotapi.KeyboardButton{{
            tgbotapi.NewKeyboardButton("Показать текущее расписание"),
        },
    }

    if role == "ADMIN" {
        adminRow := tgbotapi.NewKeyboardButtonRow(
            tgbotapi.NewKeyboardButton("Удалить кэш для Redis"),
        )
        rows = append(rows, adminRow)
    }

    keyboard := tgbotapi.NewReplyKeyboard(rows...)

    msg := tgbotapi.NewMessage(chatID, text)
    msg.ReplyMarkup = keyboard
    if _, err := bot.Send(msg); err != nil {
        log.Printf("Failed to send welcome message: %v", err)
    }
}
