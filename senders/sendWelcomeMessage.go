package senders

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func SendWelcomeMessage(bot *tgbotapi.BotAPI, chatID int64) {
	text := "Добро пожаловать! Я бот с расписанием занятий.\nНажмите кнопку ниже чтобы показать текущее расписание"
	
	keyboard := tgbotapi.NewReplyKeyboard(
        tgbotapi.NewKeyboardButtonRow(
            tgbotapi.NewKeyboardButton("Показать текущее расписание"),
        ),
    )

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = keyboard
	bot.Send(msg)
}