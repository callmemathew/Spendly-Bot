package main

import (
	"log"

	"Spendly-bot/internal/bot"
	"Spendly-bot/internal/config"
	"Spendly-bot/internal/config/expenses"
)

func main() {
	cfg := config.Load()

	if cfg.BotToken == "" {
		log.Fatal("BOT_TOKEN is empty")
	}

	expensesService := expenses.NewService()

	telegramBot, err := bot.New(cfg.BotToken, expensesService)
	if err != nil {
		log.Fatal(err)
	}

	if err := telegramBot.Run(); err != nil {
		log.Fatal(err)
	}
}
