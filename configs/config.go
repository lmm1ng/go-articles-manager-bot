package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db     DbConfig
	Bot    BotConfig
	Common CommonConfig
}

type CommonConfig struct {
	Env string
}

type DbConfig struct {
	Path string
}

type BotConfig struct {
	Token string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file. Loading default values")
	}

	return &Config{
		Db: DbConfig{
			Path: os.Getenv("DB_PATH"),
		},
		Bot: BotConfig{
			Token: os.Getenv("BOT_TOKEN"),
		},
		Common: CommonConfig{
			Env: os.Getenv("ENV"),
		},
	}
}
