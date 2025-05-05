package article

import (
	"errors"
	"fmt"
	"go-articles-manager-bot/internal/repositories/article"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

// type articleRepository interface {
// 	Create(*entities.Article) error
// }

func NewGetRandomArticleHandler(articleRepo articleRepository) th.Handler {
	return func(ctx *th.Context, update telego.Update) error {
		fmt.Println("sdasdasdad")
		var text string
		defer func() {
			ctx.Bot().
				SendMessage(
					ctx,
					tu.Message(
						update.Message.Chat.ChatID(),
						text,
					),
				)
		}()

		a, err := articleRepo.GetRandomByTgId(update.Message.From.ID)

		fmt.Println(a.Title)

		if err != nil {
			if errors.Is(err, article.ErrNotFound) {
				text = "No articles found"
				return nil
			}
		}

		text = a.ReadAt.GoString()
		return nil
	}
}
