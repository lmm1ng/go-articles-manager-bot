package article

import (
	"errors"
	"go-articles-manager-bot/internal/keyboards"
	"go-articles-manager-bot/internal/repositories/article"
	"strconv"
	"strings"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func NewDeleteArticleHandler(articleRepo articleRepository) th.Handler {
	return func(ctx *th.Context, update telego.Update) error {
		var text string

		defer func() {
			ctx.Bot().
				SendMessage(
					ctx,
					tu.Message(
						// lib is AWESOME (no)
						telego.ChatID{
							ID:       update.CallbackQuery.Message.GetChat().ID,
							Username: update.CallbackQuery.Message.GetChat().Username,
						},
						text,
					))

		}()

		articleId, err := strconv.Atoi(strings.Replace(
			update.CallbackQuery.Data,
			keyboards.DeleteArticle+" ",
			"",
			1,
		))

		if err != nil {
			text = "Article id not valid"
			return nil
		}

		if err = articleRepo.Delete(uint32(articleId)); err != nil {
			if errors.Is(err, article.ErrNotFound) {
				text = "Article not found"
			} else {
				text = "Internal error"
			}
		}

		return nil
	}
}
