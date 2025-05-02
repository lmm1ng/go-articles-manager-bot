package main

import (
	"go-articles-manager-bot/configs"
	"go-articles-manager-bot/internal/clients/telegram"
	"go-articles-manager-bot/internal/database"
	"go-articles-manager-bot/internal/handlers"
	random "go-articles-manager-bot/internal/handlers/article"
	"go-articles-manager-bot/internal/repositories/article"
	"go-articles-manager-bot/internal/repositories/user"
	"log/slog"

	th "github.com/mymmrac/telego/telegohandler"
)

func main() {
	// move to internal/logger
	logger := slog.Default()

	cfg, err := configs.LoadConfig()

	if err != nil {
		panic(err)
	}

	client, err := telegram.New(cfg.Bot.Token, logger)

	if err != nil {
		panic(err)
	}

	db, err := database.New(cfg.Db.Path, logger)

	if err != nil {
		panic(err)
	}

	userRepo := user.New(db)
	if err := userRepo.Prepare(); err != nil {
		panic(err)
	}

	aricleRepo := article.New(db)
	if err := aricleRepo.Prepare(); err != nil {
		panic(err)
	}

	randomArticleHandler := handlers.New(random.New(), th.Any())

	client.RunHandlers([]handlers.Handler{randomArticleHandler})
}
