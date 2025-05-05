package main

import (
	"go-articles-manager-bot/configs"
	"go-articles-manager-bot/internal/clients/telegram"
	"go-articles-manager-bot/internal/database"
	"go-articles-manager-bot/internal/handlers"
	"go-articles-manager-bot/internal/handlers/article"
	"go-articles-manager-bot/internal/handlers/user"
	"go-articles-manager-bot/internal/keyboards"
	"go-articles-manager-bot/internal/logger"
	"go-articles-manager-bot/internal/middlewares"
	"go-articles-manager-bot/internal/pkg/scenebuilder"

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

	createUserHandler := handlers.NewHandler(
		user.NewCreateUserHandler(userRepository),
		th.CommandEqual("start"),
	)

	createGetRandomArticleHandler := handlers.NewHandler(
		article.NewGetRandomArticleHandler(articleRepository),
		th.CommandEqual("get_random"),
	)

	createArticleScene := scenebuilder.NewScene(
		[]scenebuilder.SceneStep{
			scenebuilder.NewSceneStep(
				article.NewEnterCreateArticleHandler(),
				scenebuilder.NoScene,
			),
			scenebuilder.NewSceneStep(
				article.NewCreateArticleHandler(articleRepository, userRepository),
				scenebuilder.StepAddArticleUrl,
			),
		},
		th.TextEqual(keyboards.AddArticle),
	)

	authMiddleware := middlewares.NewAuthMiddleware()
	sceneMiddleware := middlewares.NewSceneMiddleware()

	client.
		Run(
			[]handlers.Handler{
				createUserHandler,
				createGetRandomArticleHandler,
			},
			[]th.Handler{
				sceneMiddleware,
				authMiddleware,
			},
			[]scenebuilder.Scene{
				createArticleScene,
			},
		)
}
