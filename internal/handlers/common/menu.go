package common

import (
	"go-articles-manager-bot/internal/keyboards"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

type CommonHandler struct{}

func New() *CommonHandler {
	return &CommonHandler{}
}

func (ch *CommonHandler) GetMenu() th.Handler {
	return func(ctx *th.Context, update telego.Update) error {
		ctx.Bot().SendMessage(ctx, tu.Message(update.Message.Chat.ChatID(), "Your menu is here").
			WithReplyMarkup(keyboards.NewMainMenuKeyboard()))
		return nil
	}
}
