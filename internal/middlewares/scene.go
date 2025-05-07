package middlewares

import (
	"sync"

	"go-articles-manager-bot/internal/pkg/scenebuilder"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

func NewSceneMiddleware() func(ctx *th.Context, update telego.Update) error {
	usersHash := make(map[int64]uint8)
	lock := &sync.RWMutex{}
	return func(ctx *th.Context, update telego.Update) error {

		if value := ctx.Value(scenebuilder.ScenesManagerKey); value != nil {
			return ctx.Next(update)
		}

		ctx = ctx.WithValue(
			scenebuilder.ScenesManagerKey,
			scenebuilder.NewSceneManager(&usersHash, lock),
		)

		update = update.WithContext(ctx)
		return ctx.Next(update)
	}
}
