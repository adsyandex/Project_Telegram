package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// Создаём бота
	bot, err := tgbotapi.NewBotAPI("5471771233:AAHMkVW2hezXa7xak2gVirWzhBj8XaQ_xh8")
	if err != nil {
		log.Panic(err)
	}

	// Выводим информацию об авторизации
	log.Printf("Авторизован как %s", bot.Self.UserName)

	// Удаляем существующий вебхук
	_, err = bot.Request(tgbotapi.DeleteWebhookConfig{})
	if err != nil {
		log.Fatalf("Ошибка удаления вебхука: %v", err)
	}
	log.Println("Вебхук успешно удалён.")

	// Используем метод getUpdates
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Создаём кнопки
	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("🖥️ 1 Компьютер 🛜", "https://calendar.app.google/rzNw1mfuGbxvbGj16"),
			tgbotapi.NewInlineKeyboardButtonURL("💻 2 Компьютер 🎮", "https://calendar.app.google/puEEuwoG9AFGnQPe7"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("🎰 3 Компьютер 🧩", "https://calendar.app.google/dJk6Q9d37t6zBdMu5"),
			tgbotapi.NewInlineKeyboardButtonURL("🚫 4 Компьютер 🕸️", "https://calendar.google.com/4"),
		),
	)

	// Обрабатываем входящие сообщения
	for update := range updates {
		if update.Message != nil { // Если есть сообщение
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			// Отправляем сообщение с кнопками
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите компьютер:")
			msg.ReplyMarkup = buttons

			_, err := bot.Send(msg)
			if err != nil {
				log.Printf("Ошибка отправки сообщения: %v", err)
			}
		}
	}
}
