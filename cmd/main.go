package main

import (
	"log"

	"Spendly-bot/internal/bot"
	"Spendly-bot/internal/config"
)

func main() {
	cfg := config.Load()

	if cfg.BotToken == "" {
		log.Fatal("BOT_TOKEN is empty")
	}

	telegramBot, err := bot.New(cfg.BotToken)
	if err != nil {
		log.Fatal(err)
	}

	if err := telegramBot.Run(); err != nil {
		log.Fatal(err)
	}
}
