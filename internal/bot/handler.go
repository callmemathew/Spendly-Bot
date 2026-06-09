package bot

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	stateWaitingExpense = "waiting_expense"
)

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	chatID := message.Chat.ID
	text := strings.TrimSpace(message.Text)

	if text == "/start" {
		b.userStates[chatID] = ""
		b.sendMainMenu(chatID)
		return
	}

	if text == "/menu" {
		b.userStates[chatID] = ""
		b.sendMainMenu(chatID)
		return
	}

	state := b.userStates[chatID]

	if state == stateWaitingExpense {
		b.handleExpenseInput(chatID, text)
		return
	}

	b.sendTextWithKeyboard(
		chatID,
		"Выбери действие в меню 👇",
		mainMenuKeyboard(),
	)
}

func (b *Bot) handleCallback(callback *tgbotapi.CallbackQuery) {
	if callback.Message == nil {
		return
	}

	chatID := callback.Message.Chat.ID
	messageID := callback.Message.MessageID

	switch callback.Data {
	case "menu":
		b.userStates[chatID] = ""
		b.editMainMenu(chatID, messageID)

	case "add":
		b.userStates[chatID] = stateWaitingExpense

		text := `➕ Add expense

Напиши или продиктуй расход одним сообщением.

Пример:
такси 80 овощи 100 окулист 275`

		b.editTextWithKeyboard(chatID, messageID, text, backKeyboard())

	case "stats":
		b.userStates[chatID] = ""
		b.editStats(chatID, messageID)

	case "today":
		b.userStates[chatID] = ""
		b.editToday(chatID, messageID)

	case "month":
		b.userStates[chatID] = ""
		b.editMonth(chatID, messageID)

	case "recent":
		b.userStates[chatID] = ""
		b.editRecent(chatID, messageID)

	case "settings":
		b.userStates[chatID] = ""
		b.editSettings(chatID, messageID)

	case "currency_rate":
		b.userStates[chatID] = ""

		text := `💱 Currency rate

Пока курс фиксированный:
1 EUR = 19.30 MDL

Позже подключим API и курс будет обновляться автоматически.`

		b.editTextWithKeyboard(chatID, messageID, text, backKeyboard())

	case "categories":
		b.userStates[chatID] = ""

		text := `🗂 Categories

🚕 Transport
🥦 Food
🏥 Health
🍔 Cafe
💳 Subscriptions
🏋️ Sport
🏠 Home
💸 Other`

		b.editTextWithKeyboard(chatID, messageID, text, backKeyboard())
	}

	callbackResponse := tgbotapi.NewCallback(callback.ID, "")

	_, err := b.api.Request(callbackResponse)
	if err != nil {
		log.Println("callback response error:", err)
	}
}

func (b *Bot) handleExpenseInput(chatID int64, text string) {
	if text == "" {
		b.sendTextWithKeyboard(
			chatID,
			"Не понял расход 😅\n\nПример:\nтакси 80 овощи 100",
			backKeyboard(),
		)
		return
	}

	parsedExpenses := b.expenseService.ParseText(text)

	if len(parsedExpenses) == 0 {
		b.sendTextWithKeyboard(
			chatID,
			"Не смог распознать расходы 😅\n\nПопробуй так:\nтакси 80 овощи 100 окулист 275",
			backKeyboard(),
		)
		return
	}

	var total float64
	response := "✅ Распознал расходы:\n\n"

	for _, item := range parsedExpenses {
		total += item.Amount

		response += fmt.Sprintf(
			"%s %s — %.2f MDL — %s\n",
			item.Emoji,
			item.Description,
			item.Amount,
			item.Category,
		)
	}

	response += fmt.Sprintf("\n💰 Итого: %.2f MDL", total)
	response += "\n\nПока я только распознаю расходы. Следующим этапом начну сохранять их в базу."

	b.userStates[chatID] = ""

	b.sendTextWithKeyboard(chatID, response, mainMenuKeyboard())
}

func (b *Bot) editStats(chatID int64, messageID int) {
	text := `📊 Stats

Пока статистика не подключена.

Скоро здесь будет:
🥦 Food — 25%
🚕 Transport — 15%
🏥 Health — 10%

████████░░`

	b.editTextWithKeyboard(chatID, messageID, text, statsKeyboard())
}

func (b *Bot) editToday(chatID int64, messageID int) {
	text := `📅 Today

Пока расходов за сегодня нет.

После подключения базы здесь будут траты за текущий день.`

	b.editTextWithKeyboard(chatID, messageID, text, backKeyboard())
}

func (b *Bot) editMonth(chatID int64, messageID int) {
	text := `🗓 Month

Пока месячный отчёт не подключен.

Скоро здесь будет общая сумма за месяц в MDL и EUR.`

	b.editTextWithKeyboard(chatID, messageID, text, backKeyboard())
}

func (b *Bot) editRecent(chatID int64, messageID int) {
	text := `📋 Recent

Пока последних расходов нет.

После подключения базы здесь будут последние 5–10 трат.`

	b.editTextWithKeyboard(chatID, messageID, text, backKeyboard())
}

func (b *Bot) editSettings(chatID int64, messageID int) {
	text := `⚙️ Settings

Choose what you want to configure:`

	b.editTextWithKeyboard(chatID, messageID, text, settingsKeyboard())
}

func (b *Bot) sendText(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)

	_, err := b.api.Send(msg)
	if err != nil {
		log.Println("send text error:", err)
	}
}

func (b *Bot) sendTextWithKeyboard(chatID int64, text string, keyboard tgbotapi.InlineKeyboardMarkup) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = keyboard

	_, err := b.api.Send(msg)
	if err != nil {
		log.Println("send text with keyboard error:", err)
	}
}

func (b *Bot) editTextWithKeyboard(chatID int64, messageID int, text string, keyboard tgbotapi.InlineKeyboardMarkup) {
	msg := tgbotapi.NewEditMessageTextAndMarkup(
		chatID,
		messageID,
		text,
		keyboard,
	)

	_, err := b.api.Send(msg)
	if err != nil {
		log.Println("edit message error:", err)
	}
}
