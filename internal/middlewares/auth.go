package middlewares

import (
	"fmt"
	"go-articles-manager-bot/internal/handlers"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

func NewAuthMiddleware() handlers.Cb {
	// protected := map[string]struct{}{"/addArticle": {}}

	return func(ctx *th.Context, update telego.Update) error {
		fmt.Println(update.Message)
		// if ctx.User == nil {
		// 	return errors.New("user not found")
		// }
		return ctx.Next(update)
	}
}
