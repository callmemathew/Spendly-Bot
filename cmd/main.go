package main

import (
	"log"

	_ "github.com/mattn/go-sqlite3"

	"Spendly-bot/internal/bot"
	"Spendly-bot/internal/config"
	"Spendly-bot/internal/config/expenses"
	"Spendly-bot/internal/config/storage"
)

func main() {
	cfg := config.Load()

	if cfg.BotToken == "" {
		log.Fatal("BOT_TOKEN is empty")
	}

	db, err := storage.NewSQLiteStorage("spendly.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Sqlite vse chetko!")

	expensesService := expenses.NewService(db)

	telegramBot, err := bot.New(cfg.BotToken, expensesService)
	if err != nil {
		log.Fatal(err)
	}

	if err := telegramBot.Run(); err != nil {
		log.Fatal(err)
	}
}
