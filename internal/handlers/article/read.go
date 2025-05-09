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

func (ah *ArticleHandler) NewReadArticleHandler() th.Handler {
	return func(ctx *th.Context, update telego.Update) error {
		var text string

		chatId := telego.ChatID{
			ID:       update.CallbackQuery.Message.GetChat().ID,
			Username: update.CallbackQuery.Message.GetChat().Username,
		}

		defer func() {
			ctx.Bot().SendMessage(ctx, tu.Message(chatId, text))
		}()

		args := ah.getCallbackArgs(
			update.CallbackQuery.Data,
			keyboards.ReadArticle,
			keyboards.UnreadArticle,
		)

		if len(args) != 2 {
			text = "Invalid params"
			return nil
		}
		articleId, err := strconv.Atoi(args[0])
		if err != nil {
			text = "Invalid params"
		}

		read, err := strconv.ParseBool(args[1])
		if err != nil {
			text = "Invalid params"
		}

		if err := ah.articleRepo.SetRead(uint32(articleId), !read); err != nil {
			if errors.Is(err, articleRepo.ErrNotFound) {
				text = "Article not found"
			} else {
				text = "Internal error"
			}
			return nil
		}

		ctx.Bot().DeleteMessage(
			ctx,
			&telego.DeleteMessageParams{
				ChatID:    chatId,
				MessageID: update.CallbackQuery.Message.GetMessageID(),
			},
		)

		text = "Article marked as read"

		if read {
			text = "Article marked as unread"
		}

		return nil
	}
}
