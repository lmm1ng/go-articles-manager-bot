package user

import (
	"errors"
	"go-articles-manager-bot/internal/entities"
	"go-articles-manager-bot/internal/keyboards"
	userRepo "go-articles-manager-bot/internal/repositories/user"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func (uh *UserHandler) NewCreateUserHandler() th.Handler {
	return func(ctx *th.Context, update telego.Update) error {
		var username *string
		if update.Message.From.Username != "" {
			username = &update.Message.From.Username
		}

		u := &entities.User{
			TgId:       update.Message.From.ID,
			TgUsername: username,
		}

		err := uh.userRepo.Create(u)
		if err != nil {
			text := "Internal error"
			if errors.Is(err, userRepo.ErrAlreadyExists) {
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
