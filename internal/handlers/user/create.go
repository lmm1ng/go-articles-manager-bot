package user

import (
	"errors"
	"go-articles-manager-bot/internal/entities"
	"go-articles-manager-bot/internal/keyboards"
	"go-articles-manager-bot/internal/repositories/user"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func (uh *UserHandler) NewCreateUserHandler() th.Handler {
	return func(ctx *th.Context, update telego.Update) error {
		u := &entities.User{
			TgId:       update.Message.From.ID,
			TgUsername: update.Message.From.Username,
		}

		err := uh.userRepo.Create(u)
		if err != nil {
			text := "Internal error"
			if errors.Is(err, user.ErrAlreadyExists) {
				text = "User already exists"
			}
			ctx.Bot().SendMessage(ctx, tu.Message(update.Message.Chat.ChatID(), text))
			return nil
		}

		ctx.Bot().
			SendMessage(
				ctx,
				tu.Message(
					update.Message.Chat.ChatID(),
					"Here you go",
				).
					WithReplyMarkup(keyboards.NewMainMenuKeyboard()),
			)

		return nil
	}
}
