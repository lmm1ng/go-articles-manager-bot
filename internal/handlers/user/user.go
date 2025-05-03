package user

import (
	"go-articles-manager-bot/internal/handlers"
	"go-articles-manager-bot/internal/models"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

type UserRepository interface {
	Create(user *models.User) error
}

func NewCreateUserHandler(userRepo UserRepository) handlers.Cb {
	return func(ctx *th.Context, update telego.Update) error {
		var text string
		defer ctx.Bot().SendMessage(ctx, tu.Message(update.Message.Chat.ChatID(), text))

		user := &models.User{
			TgId:       update.Message.From.ID,
			TgUsername: update.Message.From.Username,
		}

		err := userRepo.Create(user)
		if err != nil {
			text = "Error while creating user, please try again later"
			return nil
		}

		text = "Here we go"

		return nil
	}
}
