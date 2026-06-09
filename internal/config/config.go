package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BotToken string
}

func Load() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment")
	}

	return Config{
		BotToken: os.Getenv("BOT_TOKEN"),
	}
}
