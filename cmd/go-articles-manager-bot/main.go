package main

import (
	"go-articles-manager-bot/configs"
	"go-articles-manager-bot/internal/clients/telegram"
	"go-articles-manager-bot/internal/database"
	"go-articles-manager-bot/internal/handlers/article"
	"go-articles-manager-bot/internal/handlers/user"
	"go-articles-manager-bot/internal/logger"
	"go-articles-manager-bot/internal/middlewares"
	"go-articles-manager-bot/internal/models/handler"
	"go-articles-manager-bot/internal/pkg/sceneBuilder"

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

	createUserHandler := handler.New(
		user.NewCreateUserHandler(userRepository),
		th.CommandEqual("start"),
	)

	createArticleScene := sceneBuilder.NewScene(
		[]sceneBuilder.SceneStep{
			sceneBuilder.NewSceneStep(
				article.NewEnterCreateArticleHandler(),
				sceneBuilder.NoScene,
			),
			sceneBuilder.NewSceneStep(
				article.NewCreateArticleHandler(articleRepository, userRepository),
				sceneBuilder.StateAddArticleUrl,
			),
		},
		th.CommandEqual("addArticle"),
	)

	authMiddleware := middlewares.NewAuthMiddleware()
	sceneMiddleware := middlewares.NewSceneMiddleware()

	client.
		Run(
			[]handler.Handler{
				createUserHandler,
			},
			[]th.Handler{
				sceneMiddleware,
				authMiddleware,
			},
			[]sceneBuilder.Scene{
				createArticleScene,
			},
		)
}
