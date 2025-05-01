package main

import (
	"go-articles-manager-bot/configs"
	"go-articles-manager-bot/internal/clients/telegram"
	"go-articles-manager-bot/internal/handlers"
	random "go-articles-manager-bot/internal/handlers/article"
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

	randomArticleHandler := handlers.New(random.New(), th.Any())

	client.RunHandlers([]handlers.Handler{randomArticleHandler})
}
