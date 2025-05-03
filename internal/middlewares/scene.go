package middlewares

import (
	"go-articles-manager-bot/internal/models"
	"sync"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

const ScenesManagerKey = "scenesManager"

func NewSceneMiddleware() func(ctx *th.Context, update telego.Update) error {
	usersHash := make(map[int64]uint8)
	return func(ctx *th.Context, update telego.Update) error {

		if value := ctx.Value(ScenesManagerKey); value != nil {
			return ctx.Next(update)
		}

		ctx = ctx.WithValue(ScenesManagerKey, &models.ScenesManager{Users: usersHash, Mutex: sync.RWMutex{}})
		update = update.WithContext(ctx)
		return ctx.Next(update)
	}
}
