package user

import (
	"errors"
	"go-articles-manager-bot/internal/keyboards"
	"go-articles-manager-bot/internal/repositories/article"
	"strconv"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func (uh *UserHandler) NewSetUserPublicHandler() th.Handler {
	return func(ctx *th.Context, update telego.Update) error {
		var text string

		chatId := telego.ChatID{
			ID:       update.CallbackQuery.Message.GetChat().ID,
			Username: update.CallbackQuery.Message.GetChat().Username,
		}

		defer func() {
			ctx.Bot().SendMessage(ctx, tu.Message(chatId, text))
		}()

		args := uh.getCallbackArgs(
			update.CallbackQuery.Data,
			keyboards.SetPublic,
			keyboards.HideUser,
		)

		if len(args) != 1 {
			text = "Invalid params"
			return nil
		}

		public, err := strconv.ParseBool(args[0])
		if err != nil {
			text = "Invalid params"
		}

		if err := uh.userRepo.UpdatePublicByTgId(update.CallbackQuery.From.ID, !public); err != nil {
			if errors.Is(err, article.ErrNotFound) {
				text = "User not found"
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

		text = "Your account is public for now"

		if public {
			text = "Your account is hidden for now"
		}

		return nil
	}
}
