package controllers

import (
    "log"
    "net/http"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type WebhookController struct {
    Bot     *tgbotapi.BotAPI
    Updates tgbotapi.UpdatesChannel
}

func (wc *WebhookController) HandleWebhook(w http.ResponseWriter, r *http.Request) {
    for update := range wc.Updates {
        if update.Message != nil {
            log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

            // buttons Создаём кнопки
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

            // Отправляем сообщение с кнопками
            msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите компьютер:")
            msg.ReplyMarkup = buttons

            _, err := wc.Bot.Send(msg)
            if err != nil {
                log.Printf("Ошибка отправки сообщения: %v", err)
            }
        }
    }
}
