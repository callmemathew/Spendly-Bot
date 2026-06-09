package bot

import (
	"log"

	"Spendly-bot/internal/config/expenses"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	api            *tgbotapi.BotAPI
	expenseService *expenses.Service
	userStates     map[int64]string
}

func New(token string, expenseService *expenses.Service) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &Bot{
		api:            api,
		expenseService: expenseService,
		userStates:     make(map[int64]string),
	}, nil
}

func (b *Bot) Run() error {
	log.Println("Твой ботяра:", b.api.Self.UserName, "запущен")

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := b.api.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message != nil {
			b.handleMessage(update.Message)
		}

		if update.CallbackQuery != nil {
			b.handleCallback(update.CallbackQuery)
		}
	}

	return nil
}
