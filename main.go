package main

import (
	"Project_Telegram/quickstart/controllers"
	"log"
	"net/http"

	"github.com/astaxie/beego"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)
   // main функция подключения IP бота
func main() {
	bot, err := tgbotapi.NewBotAPI("5471771233:AAHMkVW2hezXa7xak2gVirWzhBj8XaQ_xh8")
	if err != nil {
		log.Panic(err)
	}

	updates := bot.ListenForWebhook("/webhook")
	http.HandleFunc("/telegram_webhook", func(w http.ResponseWriter, r *http.Request) {
		controller := &controllers.WebhookController{Bot: bot, Updates: updates}
		controller.HandleWebhook(w, r)
	})

	beego.BConfig.Listen.HTTPSCertFile = "path/to/cert.pem"
	beego.BConfig.Listen.HTTPSKeyFile = "path/to/key.pem"
	beego.BConfig.Listen.HTTPPort = 443
	beego.BConfig.Listen.EnableHTTP = false
	beego.BConfig.Listen.EnableHTTPS = true

	log.Println("Запуск сервера...")
	beego.Run()
}
