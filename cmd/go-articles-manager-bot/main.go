package main

import (
	"go-articles-manager-bot/configs"
	"go-articles-manager-bot/internal/clients/telegram"
	"go-articles-manager-bot/internal/database"
	"go-articles-manager-bot/internal/handlers"
	articleHandler "go-articles-manager-bot/internal/handlers/article"
	userHandler "go-articles-manager-bot/internal/handlers/user"
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

	articleHandler := articleHandler.New(articleRepository, userRepository)

	// Article section

	createGetRandomArticleHandler := handlers.NewHandler(
		articleHandler.NewGetRandomArticleHandler(),
		th.TextEqual(keyboards.RandomArticle),
	)

	createGetArticleByIdHandler := handlers.NewHandler(
		articleHandler.NewGetArticleByIdHandler(),
		th.CallbackDataPrefix(keyboards.SelectArticle),
	)

	readArticleHanler := handlers.NewHandler(
		articleHandler.NewReadArticleHandler(),
		th.Or(th.CallbackDataPrefix(keyboards.ReadArticle), th.CallbackDataPrefix(keyboards.UnreadArticle)),
	)

	deleteArticleHanler := handlers.NewHandler(
		articleHandler.NewDeleteArticleHandler(),
		th.CallbackDataPrefix(keyboards.DeleteArticle),
	)

	showArticlesHandler := handlers.NewHandler(articleHandler.NewShowArticlesHandler(),
		th.TextEqual(keyboards.ShowArticles),
	)

	showArticlesChangePageHandler := handlers.NewHandler(
		articleHandler.NewShowArticlesChangePageHandler(),
		th.Or(
			th.CallbackDataPrefix(keyboards.PrevPage),
			th.CallbackDataPrefix(keyboards.NextPage),
		),
	)

	showArticlesChangeVisibilityHandler := handlers.NewHandler(articleHandler.NewShowArticlesChangeVisibilityHandler(),
		th.Or(
			th.CallbackDataPrefix(keyboards.HideRead),
			th.CallbackDataPrefix(keyboards.ShowRead),
		))

	createArticleScene := scenebuilder.NewScene(
		[]scenebuilder.SceneStep{
			scenebuilder.NewSceneStep(
				articleHandler.NewEnterCreateArticleHandler(),
				scenebuilder.NoScene,
			),
			scenebuilder.NewSceneStep(
				articleHandler.NewCreateArticleHandler(),
				scenebuilder.StepAddArticleUrl,
			),
		},
		th.TextEqual(keyboards.AddArticle),
	)

	// User section

	userHandler := userHandler.New(userRepository)

	createUserHandler := handlers.NewHandler(
		userHandler.NewCreateUserHandler(),
		th.CommandEqual("start"),
	)

	// Middlewares

	authMiddleware := middlewares.NewAuthMiddleware()
	sceneMiddleware := middlewares.NewSceneMiddleware()

	client.
		Run(
			[]handlers.Handler{
				createUserHandler,
				createGetRandomArticleHandler,
				readArticleHanler,
				deleteArticleHanler,
				showArticlesHandler,
				showArticlesChangePageHandler,
				showArticlesChangeVisibilityHandler,
				createGetArticleByIdHandler,
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
