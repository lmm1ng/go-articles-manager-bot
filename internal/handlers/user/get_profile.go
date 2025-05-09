package user

import (
	"errors"
	"fmt"
	"go-articles-manager-bot/internal/keyboards"
	userRepo "go-articles-manager-bot/internal/repositories/user"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func (uh *UserHandler) NewGetUserProfileHandler() th.Handler {
	return func(ctx *th.Context, update telego.Update) error {
		u, err := uh.userRepo.GetByTgId(update.Message.From.ID)

		if err != nil {
			if errors.Is(err, userRepo.ErrNotFound) {
				ctx.Bot().SendMessage(ctx, tu.Message(update.Message.Chat.ChatID(), "User not found"))
			} else {
				ctx.Bot().SendMessage(ctx, tu.Message(update.Message.Chat.ChatID(), "Internal error"))
			}
			return nil
		}

		var desc string

		if u.Desc == nil {
			desc = "No bio"
		} else {
			desc = *u.Desc
		}

		var username string

		if u.TgUsername == nil {
			username = "Hidden"
		} else {
			username = fmt.Sprintf("@%s", *u.TgUsername)
		}

		ctx.Bot().SendMessage(
			ctx,
			tu.Message(
				update.Message.Chat.ChatID(),
				fmt.Sprintf("%s\n\n%s", username, desc),
			).
				WithReplyMarkup(keyboards.NewProfileInlineKeyboard(!u.Public)),
		)

		return nil
	}
}
