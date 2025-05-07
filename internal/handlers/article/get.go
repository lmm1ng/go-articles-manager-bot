package article

import (
	"errors"
	"fmt"
	"go-articles-manager-bot/internal/keyboards"
	"go-articles-manager-bot/internal/repositories/article"
	"strconv"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

// type articleRepository interface {
// 	Create(*entities.Article) error
// }

func (ah *ArticleHandler) NewGetRandomArticleHandler() th.Handler {
	return func(ctx *th.Context, update telego.Update) error {
		var text string

		a, err := ah.articleRepo.GetRandomByTgId(update.Message.From.ID)
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

		ctx.Bot().
			SendMessage(
				ctx,
				tu.Message(
					update.Message.Chat.ChatID(),
					a.GetTitleLink(),
				).
					WithReplyMarkup(keyboards.NewArticleInlineKeyboard(a.Id, true)).
					WithParseMode("Markdown"),
			)
		return nil
	}
}

func (ah *ArticleHandler) NewGetArticleByPosHandler() th.Handler {
	return func(ctx *th.Context, update telego.Update) error {
		chatId := telego.ChatID{
			ID:       update.CallbackQuery.Message.GetChat().ID,
			Username: update.CallbackQuery.Message.GetChat().Username,
		}

		args := ah.getCallbackArgs(update.CallbackQuery.Data, keyboards.SelectArticle)

		if len(args) != 3 {
			ctx.Bot().SendMessage(ctx, tu.Message(chatId, "Invalid params"))
			return nil
		}

		pos, err := strconv.ParseBool(args[0])

		if err != nil {
			ctx.Bot().SendMessage(ctx, tu.Message(chatId, "Invalid params"))
			return nil
		}

		a, err := ah.articleRepo.GetById(update.Message.From.ID)
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
