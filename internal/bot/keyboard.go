package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (b *Bot) sendMainMenu(chatID int64) {
	text := `💸 Spendly

Your smart Telegram expense tracker.

Add expenses by text or dictation, check category stats, and keep your budget under control in MDL and EUR.`

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("➕ Add expense", "add"),
			tgbotapi.NewInlineKeyboardButtonData("📊 Stats", "stats"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("📅 Today", "today"),
			tgbotapi.NewInlineKeyboardButtonData("🗓 Month", "month"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("📋 Recent", "recent"),
			tgbotapi.NewInlineKeyboardButtonData("⚙️ Settings", "settings"),
		),
	)

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = keyboard

	b.api.Send(msg)
}
