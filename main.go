package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// Укажите токен вашего бота
	bot, err := tgbotapi.NewBotAPI("5471771233:AAHMkVW2hezXa7xak2gVirWzhBj8XaQ_xh8")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Авторизован как %s", bot.Self.UserName)

	// Обработка /start команды
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil && update.Message.Text == "/start" {
			// Создание кнопок
			button1 := tgbotapi.NewInlineKeyboardButtonURL("1 Комп", "https://calendar.google.com/calendar/u/0/r/eventedit?text=Бронь+1+Комп")
			button2 := tgbotapi.NewInlineKeyboardButtonURL("2 Комп", "https://calendar.google.com/calendar/u/0/r/eventedit?text=Бронь+2+Комп")
			button3 := tgbotapi.NewInlineKeyboardButtonURL("3 Комп", "https://calendar.google.com/calendar/u/0/r/eventedit?text=Бронь+3+Комп")
      
			// Создание клавиатуры
			keyboard := tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(button1),
				tgbotapi.NewInlineKeyboardRow(button2),
				tgbotapi.NewInlineKeyboardRow(button3),
			)

			// Отправка сообщения с кнопками
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите компьютер для бронирования:")
			msg.ReplyMarkup = keyboard

			_, err := bot.Send(msg)
			if err != nil {
				log.Panic(err)
			}
		}
	}
}

