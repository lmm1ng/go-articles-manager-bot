package article

import (
	"go-articles-manager-bot/internal/keyboards"
	"strconv"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func (ah *ArticleHandler) NewShowArticlesHandler() th.Handler {
	return func(ctx *th.Context, update telego.Update) error {
		articles, err := ah.getArticlesList(update.Message.From.ID, 1, true)
		if err != nil {
			ctx.Bot().
				SendMessage(ctx, tu.Message(update.Message.Chat.ChatID(), err.Error()))
		}

		ctx.Bot().
			SendMessage(
				ctx,
				tu.Message(
					update.Message.Chat.ChatID(),
					ah.generateAtriclesMessage(articles),
				).
					WithParseMode("Markdown").
					WithLinkPreviewOptions(&telego.LinkPreviewOptions{IsDisabled: true}).
					WithReplyMarkup(keyboards.NewArticlesListInlineKeyboard(1, true, articles)),
			)

		return nil
	}
}

func (ah *ArticleHandler) NewShowArticlesChangePageHandler() th.Handler {
	return func(ctx *th.Context, update telego.Update) error {
		var text string

		args := ah.getCallbackArgs(update.CallbackQuery.Data, keyboards.PrevPage, keyboards.NextPage)

		if len(args) != 2 {
			text = "Invalid params"
		}

		page, err := strconv.Atoi(args[0])
		if err != nil {
			text = "Invalid params"
		}

		read, err := strconv.ParseBool(args[1])
		if err != nil {
			text = "Invalid params"
		}

		articles, err := ah.getArticlesList(
			update.CallbackQuery.Message.GetChat().ID,
			uint16(page),
			read,
		)

		if err != nil {
			text = "Articles not found"
		}

		chatId := telego.ChatID{
			ID:       update.CallbackQuery.Message.GetChat().ID,
			Username: update.CallbackQuery.Message.GetChat().Username,
		}

		if text != "" {
			ctx.Bot().SendMessage(ctx, tu.Message(chatId, text))
			return nil
		}

		ctx.Bot().DeleteMessage(ctx, &telego.DeleteMessageParams{
			ChatID:    chatId,
			MessageID: update.CallbackQuery.Message.GetMessageID(),
		})

		ctx.Bot().
			SendMessage(
				ctx,
				tu.Message(
					chatId,
					ah.generateAtriclesMessage(articles),
				).
					WithParseMode("Markdown").
					WithLinkPreviewOptions(&telego.LinkPreviewOptions{IsDisabled: true}).
					WithReplyMarkup(keyboards.NewArticlesListInlineKeyboard(uint16(page), read, articles)),
			)
		return nil
	}
}

func (ah *ArticleHandler) NewShowArticlesChangeVisibilityHandler() th.Handler {
	return func(ctx *th.Context, update telego.Update) error {
		chatId := telego.ChatID{
			ID:       update.CallbackQuery.Message.GetChat().ID,
			Username: update.CallbackQuery.Message.GetChat().Username,
		}

		args := ah.getCallbackArgs(update.CallbackQuery.Data, keyboards.HideRead, keyboards.ShowRead)

		if len(args) != 1 {
			ctx.Bot().SendMessage(ctx, tu.Message(chatId, "Invalid params"))
			return nil
		}

		read, err := strconv.ParseBool(args[0])

		if err != nil {
			ctx.Bot().SendMessage(ctx, tu.Message(chatId, "Invalid params"))
			return nil
		}

		articles, err := ah.getArticlesList(
			update.CallbackQuery.Message.GetChat().ID,
			1,
			read,
		)

		if err != nil {
			ctx.Bot().SendMessage(ctx, tu.Message(chatId, "Articles not found"))
			return nil
		}

		ctx.Bot().DeleteMessage(
			ctx,
			&telego.DeleteMessageParams{
				ChatID:    chatId,
				MessageID: update.CallbackQuery.Message.GetMessageID(),
			},
		)

		ctx.Bot().
			SendMessage(
				ctx,
				tu.Message(
					chatId,
					ah.generateAtriclesMessage(articles),
				).
					WithParseMode("Markdown").
					WithLinkPreviewOptions(&telego.LinkPreviewOptions{IsDisabled: true}).
					WithReplyMarkup(keyboards.NewArticlesListInlineKeyboard(1, read, articles)),
			)
		return nil
	}
}
