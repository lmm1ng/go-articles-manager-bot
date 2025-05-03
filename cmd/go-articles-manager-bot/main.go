package main

import (
	"go-articles-manager-bot/configs"
	"go-articles-manager-bot/internal/clients/telegram"
	"go-articles-manager-bot/internal/database"
	"go-articles-manager-bot/internal/handlers"
	"go-articles-manager-bot/internal/handlers/article"
	"go-articles-manager-bot/internal/handlers/user"
	"go-articles-manager-bot/internal/logger"
	"go-articles-manager-bot/internal/middlewares"

	articleRepo "go-articles-manager-bot/internal/repositories/article"
	userRepo "go-articles-manager-bot/internal/repositories/user"

	th "github.com/mymmrac/telego/telegohandler"
)

func main() {
	cfg := configs.Load()

	log := logger.New(cfg.Common.Env)

	client := telegram.MustNew(cfg.Bot.Token, log)

	db := database.MustNew(cfg.Db.Path, log)

	userRepository := userRepo.New(db)
	articleRepository := articleRepo.New(db)

	createUserHandler := handlers.New(user.NewCreateUserHandler(userRepository), th.CommandEqual("start"))
	enterCreareArticleHandler := handlers.New(article.NewEnterCreateArticleHandler(), th.CommandEqual("add"))
	createArticleHandler := handlers.New(article.NewCreateArticleHandler(articleRepository, userRepository), th.TextPrefix("http"))

	authMiddleware := middlewares.NewAuthMiddleware()
	sceneMiddleware := middlewares.NewSceneMiddleware()

	client.RunHandlers([]handlers.Handler{createUserHandler, enterCreareArticleHandler, createArticleHandler}, []handlers.Cb{sceneMiddleware, authMiddleware})
}
