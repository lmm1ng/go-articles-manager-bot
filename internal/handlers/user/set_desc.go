package user

import (
	"errors"
	userRepo "go-articles-manager-bot/internal/repositories/user"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func (uh *UserHandler) NewEnterSetUserDescHandler() th.Handler {
	return func(ctx *th.Context, update telego.Update) error {
		chatId := telego.ChatID{
			ID:       update.CallbackQuery.Message.GetChat().ID,
			Username: update.CallbackQuery.Message.GetChat().Username,
		}

		ctx.Bot().SendMessage(ctx, tu.Message(chatId, "Write a few words about yourself:"))
		return nil
	}
}

func (uh *UserHandler) NewSetUserDescHandler() th.Handler {
	return func(ctx *th.Context, update telego.Update) error {
		var text string
		err := uh.userRepo.UpdateDescByTgId(update.Message.From.ID, update.Message.Text)

		if err != nil {
			if errors.Is(err, userRepo.ErrNotFound) {
				text = "Unknown user"
			}
			text = "Internal error"
		} else {
			text = "Profile description successfully set"
		}

		ctx.Bot().SendMessage(ctx, tu.Message(update.Message.Chat.ChatID(), text))
		return nil
	}
}
