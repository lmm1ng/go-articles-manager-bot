package middlewares

import (
	"go-articles-manager-bot/internal/pkg/scenebuilder"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

func NewSceneMiddleware() func(ctx *th.Context, update telego.Update) error {
	manager := scenebuilder.NewSceneManager()
	return func(ctx *th.Context, update telego.Update) error {
		ctx = ctx.WithValue(
			scenebuilder.ScenesManagerKey,
			manager,
		)

		update = update.WithContext(ctx)
		return ctx.Next(update)
	}
}
