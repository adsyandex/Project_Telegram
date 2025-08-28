package main

import (
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Глобальная переменная для бота (для простоты примера)
var bot *tgbotapi.BotAPI

func main() {
	// 1. Загрузка конфигурации
	godotenv.Load()
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Panic("BOT_TOKEN is not set")
	}

	// 2. Инициализация бота
	var err error
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// 3. Настройка Webhook URL
	// ВАЖНО: ЗАМЕНИТЕ "YOUR_APP_URL" на реальный URL вашего приложения на TimeWeb!
	// Например: "https://my-cool-bot.timeweb.cloud"
	/*webhookURL := os.Getenv("WEBHOOK_URL") + "/webhook" // Будет обрабатывать запросы на /webhook
	wh, err := tgbotapi.NewWebhook(webhookURL)
	if err != nil {
		log.Fatal(err)
	}*/
    webhookURL := os.Getenv("WEBHOOK_URL")
    if webhookURL == "" {
       log.Println("WARNING: WEBHOOK_URL is not set. Webhook will not be set. Waiting for configuration...")
    // Можно просто запустить сервер без установки вебхука, чтобы получить URL
    // Или завершить работу с ошибкой, в зависимости от логики
    } else {
       wh, err := tgbotapi.NewWebhook(webhookURL + "/webhook")
       if err != nil {
           log.Fatal(err)
    }

	/* 4. Установка Webhook'а у Telegram
	_, err = bot.Request(wh)
	if err != nil {
		log.Fatal(err)
	}*/
    // 4. Установка Webhook'а у Telegram
    _, err = bot.Request(wh)
    if err != nil {
        log.Fatal(err)  // Было log.Fatal(app) - это ошибка!
    }
    log.Printf("Webhook set to: %s", webhookURL)
}

	// 5. Настройка Gin HTTP-сервера
	router := gin.Default()

	// Отдаем какую-нибудь заглушку на главной странице, чтобы хостинг видел, что приложение живо
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Telegram Bot is running!"})
	})

	// Это основной endpoint, куда Telegram будет присылать обновления
	router.POST("/webhook", func(c *gin.Context) {
		// Gin автоматически парсит JSON из тела запроса в структуру Update
		var update tgbotapi.Update
		if err := c.BindJSON(&update); err != nil {
			log.Printf("Error parsing update: %v", err)
			c.Status(http.StatusBadRequest)
			return
		}

		// Запускаем обработку апдейта в горутине, чтобы быстро ответить Telegram
		go handleUpdate(update)
		c.Status(http.StatusOK)
	})

	// 6. Запуск HTTP-сервера
	// Хостинг сам передаст нужный порт (часто через переменную окружения PORT)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Порт по умолчанию для локальной разработки
	}
	router.Run(":" + port)
}

func handleUpdate(update tgbotapi.Update) {
	// Ваша логика обработки сообщений переносится сюда
	if update.Message != nil {
		handleMessage(update.Message)
	}
	// Здесь же можно будет обрабатывать CallbackQuery от кнопок
}

func handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	// Создаём кнопки (как у вас было)
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

	msg := tgbotapi.NewMessage(message.Chat.ID, "Выберите компьютер:")
	msg.ReplyMarkup = buttons

	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}