package bot

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	chatID := message.Chat.ID
	text := strings.TrimSpace(message.Text)

	if text == "/start" {
		b.sendMainMenu(chatID)
		return
	}

	response := "✅ Я получил сообщение:\n\n" + text + "\n\nСкоро я научусь сохранять расходы."

	b.sendText(chatID, response)
}

func (b *Bot) handleCallback(callback *tgbotapi.CallbackQuery) {
	chatID := callback.Message.Chat.ID

	switch callback.Data {
	case "add":
		b.sendText(chatID, "➕ Напиши или продиктуй расход:\n\nтакси 80 овощи 100 окулист 275")

	case "stats":
		b.sendText(chatID, "📊 Статистика скоро будет здесь.")

	case "today":
		b.sendText(chatID, "📅 Сегодня расходов пока нет.")

	case "month":
		b.sendText(chatID, "🗓 За месяц расходов пока нет.")

	case "recent":
		b.sendText(chatID, "📋 Последние расходы скоро будут здесь.")

	case "settings":
		b.sendText(chatID, "⚙️ Настройки скоро будет здесь.")
	}

	callbackResponse := tgbotapi.NewCallback(callback.ID, "")

	_, err := b.api.Request(callbackResponse)
	if err != nil {
		log.Println("callback response error:", err)
	}
}

func (b *Bot) sendText(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)

	_, err := b.api.Send(msg)
	if err != nil {
		log.Println("send text error:", err)
	}
}
