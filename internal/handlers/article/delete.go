package article

import (
	"errors"
	"go-articles-manager-bot/internal/keyboards"
	articleRepo "go-articles-manager-bot/internal/repositories/article"
	"strconv"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func (ah *ArticleHandler) NewDeleteArticleHandler() th.Handler {
	return func(ctx *th.Context, update telego.Update) error {
		chatId := telego.ChatID{
			ID:       update.CallbackQuery.Message.GetChat().ID,
			Username: update.CallbackQuery.Message.GetChat().Username,
		}

		args := ah.getCallbackArgs(update.CallbackQuery.Data, keyboards.DeleteArticle)

		if len(args) != 1 {
			ctx.Bot().SendMessage(ctx, tu.Message(chatId, "Invalid params"))
			return nil
		}

		articleId, err := strconv.Atoi(args[0])
		if err != nil {
			ctx.Bot().SendMessage(ctx, tu.Message(chatId, "Invalid params"))
			return nil
		}

		if err = ah.articleRepo.Delete(uint32(articleId)); err != nil {
			var errText string
			if errors.Is(err, articleRepo.ErrNotFound) {
				errText = "Article not found"
			} else {
				errText = "Internal error"
			}

			ctx.Bot().SendMessage(ctx, tu.Message(chatId, errText))
			return nil
		}

		ctx.Bot().DeleteMessage(
			ctx,
			&telego.DeleteMessageParams{
				ChatID:    chatId,
				MessageID: update.CallbackQuery.Message.GetMessageID(),
			},
		)

		ctx.Bot().SendMessage(ctx, tu.Message(chatId, "Article deleted"))

		return nil
	}
}
