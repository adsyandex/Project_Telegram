package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Загрузка конфигурации
	godotenv.Load()
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Panic("BOT_TOKEN is not set")
	}

	// Инициализация бота
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Авторизован как %s", bot.Self.UserName)

	// Удаляем вебхук для чистого режима GetUpdates
	_, err = bot.Request(tgbotapi.DeleteWebhookConfig{})
	if err != nil {
		log.Printf("Предупреждение при удалении вебхука: %v", err)
	}

	// Настраиваем Long Polling
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	// Создаём кнопки
	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("🖥️ 1 Компьютер ☕🍔", "https://calendar.app.google/rzNw1mfuGbxvbGj16"),
			tgbotapi.NewInlineKeyboardButtonURL("💻 2 Компьютер 🍸🍰", "https://calendar.app.google/puEEuwoG9AFGnQPe7"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("🎰 3 Компьютер 🍹🌭", "https://calendar.app.google/dJk6Q9d37t6zBdMu5"),
			tgbotapi.NewInlineKeyboardButtonURL("🕹️ 4 Компьютер 🧃🥪", "https://calendar.app.google/..."),
		),
	)

	// Обрабатываем входящие сообщения
	log.Println("Бот запущен в режиме Long Polling...")
	for update := range updates {
		if update.Message != nil {
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
