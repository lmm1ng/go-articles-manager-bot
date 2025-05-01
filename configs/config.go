package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db  DbConfig
	Bot BotConfig
}

type DbConfig struct {
	Path string
}

type BotConfig struct {
	Token string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("Error loading .env file, %w", err)
	}

	return &Config{
		Db: DbConfig{
			Path: os.Getenv("DB_PATH"),
		},
		Bot: BotConfig{
			Token: os.Getenv("BOT_TOKEN"),
		},
	}, nil
}
