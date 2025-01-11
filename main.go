package main

import (
	"log"
	"net/http"

	"github.com/astaxie/beego"
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

	// Указываем публичный URL для вебхука (замените YOUR_PUBLIC_URL на реальный URL)
	webhookURL := "https://adsyandex-project-telegram-041c.twc1.net:8081"

	// Устанавливаем вебхук
	webhookConfig, err := tgbotapi.NewWebhookWithCert(webhookURL, nil)
	if err != nil {
		log.Fatalf("Ошибка создания конфигурации вебхука: %v", err)
	}
	_, err = bot.Request(webhookConfig)
	if err != nil {
		log.Fatalf("Ошибка установки вебхука: %v", err)
	}

	// Подтверждаем, что вебхук установлен
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatalf("Ошибка получения информации о вебхуке: %v", err)
	}
	if info.URL != webhookURL {
		log.Fatalf("Установлен неправильный вебхук: %s", info.URL)
	}

	// Обрабатываем обновления через вебхук
	updates := bot.ListenForWebhook("/")

	// Настраиваем сервер на порту 8081
	go func() {
		//log.Fatal(http.ListenAndServe(":8081", nil))
		//}()
		log.Fatal(http.ListenAndServeTLS(":443", "path/to/cert.pem", "path/to/key.pem", nil))
	}()

	// Создаём кнопки
	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("🖥️ 1 Компьютер ☕🍔", "https://calendar.app.google/rzNw1mfuGbxvbGj16"),
			tgbotapi.NewInlineKeyboardButtonURL("💻 2 Компьютер 🍸🍰", "https://calendar.app.google/puEEuwoG9AFGnQPe7"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("🎰 3 Компьютер 🍹🌭", "https://calendar.app.google/dJk6Q9d37t6zBdMu5"),
			tgbotapi.NewInlineKeyboardButtonURL("🚫 4 Компьютер 🧃🥪", "https://calendar.google.com/4"),
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

	// Настраиваем Beego для запуска HTTP-сервера
	beego.Run()
}
