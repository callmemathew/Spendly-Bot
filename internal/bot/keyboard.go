package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func mainMenuKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
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
}

func backKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬅️ Back to menu", "menu"),
		),
	)
}

func statsKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("📅 Today", "today"),
			tgbotapi.NewInlineKeyboardButtonData("🗓 Month", "month"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("📋 Recent", "recent"),
			tgbotapi.NewInlineKeyboardButtonData("⬅️ Back", "menu"),
		),
	)
}

func settingsKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("💱 Currency rate", "currency_rate"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🗂 Categories", "categories"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("⬅️ Back", "menu"),
		),
	)
}

func mainMenuText() string {
	return `💸 Spendly

Your smart Telegram expense tracker.

Add expenses by text or dictation, check category stats, and keep your budget under control in MDL and EUR.`
}

func (b *Bot) sendMainMenu(chatID int64) {
	b.sendTextWithKeyboard(chatID, mainMenuText(), mainMenuKeyboard())
}

func (b *Bot) editMainMenu(chatID int64, messageID int) {
	b.editTextWithKeyboard(chatID, messageID, mainMenuText(), mainMenuKeyboard())
}
