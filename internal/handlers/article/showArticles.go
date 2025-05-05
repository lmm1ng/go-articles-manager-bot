package article

package article

import (
	"errors"
	"fmt"
	"go-articles-manager-bot/internal/keyboards"
	"go-articles-manager-bot/internal/repositories/article"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

// type articleRepository interface {
// 	Create(*entities.Article) error
// }

func NewShowArticlesHandler(articleRepo articleRepository) th.Handler {
	return func(ctx *th.Context, update telego.Update) error {
		var text string

		a, err := articleRepo.GetRandomByTgId(update.Message.From.ID)
		if err != nil {
			if errors.Is(err, article.ErrNotFound) {
				text = "No articles found"
			} else {
				text = "Internal error"
			}

			ctx.Bot().
				SendMessage(
					ctx,
					tu.Message(
						update.Message.Chat.ChatID(),
						text,
					),
				)
			return nil
		}

		if a.Title != nil {
			text = fmt.Sprintf("[%s](%s)", *a.Title, a.Url)
		} else {
			text = a.Url
		}

		ctx.Bot().
			SendMessage(
				ctx,
				tu.Message(
					update.Message.Chat.ChatID(),
					text,
				).
					WithReplyMarkup(keyboards.NewArticleInlineKeyboard(a.Id, true)).
					WithParseMode("Markdown"),
			)
		return nil
	}
}
