package article

import (
	"errors"
	"fmt"
	"go-articles-manager-bot/internal/keyboards"
	"go-articles-manager-bot/internal/repositories/article"
	articleRepo "go-articles-manager-bot/internal/repositories/article"
	"strconv"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func (ah *ArticleHandler) NewGetRandomArticleHandler() th.Handler {
	return func(ctx *th.Context, update telego.Update) error {
		a, err := ah.articleRepo.GetRandomByTgId(update.Message.From.ID)
		if err != nil {
			var errText string
			if errors.Is(err, articleRepo.ErrNotFound) {
				errText = "No articles found"
			} else {
				errText = "Internal error"
			}

			ctx.Bot().SendMessage(ctx, tu.Message(update.Message.Chat.ChatID(), errText))
			return nil
		}

		ctx.Bot().
			SendMessage(
				ctx,
				tu.Message(
					update.Message.Chat.ChatID(),
					a.GetTitleLink(),
				).
					WithReplyMarkup(keyboards.NewArticleInlineKeyboard(a.Id, a.ReadAt != nil, true)).
					WithParseMode("Markdown"),
			)
		return nil
	}
}

func (ah *ArticleHandler) NewGetArticleByIdHandler() th.Handler {
	return func(ctx *th.Context, update telego.Update) error {
		chatId := telego.ChatID{
			ID:       update.CallbackQuery.Message.GetChat().ID,
			Username: update.CallbackQuery.Message.GetChat().Username,
		}

		args := ah.getCallbackArgs(update.CallbackQuery.Data, keyboards.SelectArticle)

		if len(args) != 1 {
			ctx.Bot().SendMessage(ctx, tu.Message(chatId, "Invalid params"))
			return nil
		}

		id, err := strconv.Atoi(args[0])

		if err != nil {
			ctx.Bot().SendMessage(ctx, tu.Message(chatId, "Invalid params"))
			return nil
		}

		a, err := ah.articleRepo.GetById(uint32(id))
		if err != nil {
			errText := "Internal error"
			if errors.Is(err, article.ErrNotFound) {
				errText = "No articles found"
			}

			ctx.Bot().SendMessage(ctx, tu.Message(chatId, errText))
			return nil
		}

		ctx.Bot().
			SendMessage(
				ctx,
				tu.Message(
					chatId,
					a.GetTitleLink(),
				).
					WithReplyMarkup(keyboards.NewArticleInlineKeyboard(a.Id, a.ReadAt != nil, true)).
					WithParseMode("Markdown"),
			)
		return nil
	}
}

func (ah *ArticleHandler) NewGetVibeArticleHandler() th.Handler {
	return func(ctx *th.Context, update telego.Update) error {
		a, err := ah.articleRepo.GetVibe(update.Message.From.ID)
		if err != nil {
			var text string
			if errors.Is(err, articleRepo.ErrNotFound) {
				text = "No articles found"
			} else {
				text = "Internal error"
			}

			ctx.Bot().SendMessage(ctx, tu.Message(update.Message.Chat.ChatID(), text))
			return nil
		}

		u, err := ah.userRepo.GetById(a.UserId)
		if err != nil {
			ctx.Bot().SendMessage(ctx, tu.Message(update.Message.Chat.ChatID(), "Internal error"))
			return nil
		}

		username := "hidden one"

		if u.TgUsername != nil {
			username = *u.TgUsername
		}

		var text string

		if u.Desc != nil {
			text = fmt.Sprintf("Article of @%s\n(%s)\n\n%s", username, *u.Desc, a.GetTitleLink())

		} else {
			text = fmt.Sprintf("Article of @%s\n\n%s", username, a.GetTitleLink())
		}

		ctx.Bot().
			SendMessage(
				ctx,
				tu.Message(
					update.Message.Chat.ChatID(),
					text,
				).
					WithParseMode("Markdown"),
			)
		return nil
	}
}
