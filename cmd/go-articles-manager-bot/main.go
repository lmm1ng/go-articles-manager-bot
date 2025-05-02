package main

import (
	"go-articles-manager-bot/configs"
	"go-articles-manager-bot/internal/clients/telegram"
	"go-articles-manager-bot/internal/database"
	"go-articles-manager-bot/internal/handlers"
	random "go-articles-manager-bot/internal/handlers/article"
	"go-articles-manager-bot/internal/logger"
	"go-articles-manager-bot/internal/repositories/article"
	"go-articles-manager-bot/internal/repositories/user"

	th "github.com/mymmrac/telego/telegohandler"
)

func main() {
	cfg := configs.Load()

	log := logger.New(cfg.Common.Env)

	client := telegram.MustNew(cfg.Bot.Token, log)

	db := database.MustNew(cfg.Db.Path, log)

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
