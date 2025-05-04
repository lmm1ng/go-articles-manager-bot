package middlewares

import (
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

func NewAuthMiddleware() th.Handler {
	// protected := map[string]struct{}{"/addArticle": {}}

	return func(ctx *th.Context, update telego.Update) error {
		// if ctx.User == nil {
		// 	return errors.New("user not found")
		// }
		return ctx.Next(update)
	}
}
