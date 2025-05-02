package article

import (
	"fmt"
	"go-articles-manager-bot/internal/handlers"
	"sync"

	"go-articles-manager-bot/internal/repositories/article"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func NewRandomArticleHandler(articleRepo *article.Repository) handlers.Cb {
	return func(ctx *th.Context, update telego.Update) error {
		ctx.Bot().SendMessage(ctx, tu.Message(
			tu.ID(update.Message.Chat.ID),
			fmt.Sprintf("Hello %s!", update.Message.From.FirstName),
		))
		return nil
	}
}

func NewCreateArticleHandler(articleRepo *article.Repository) handlers.Cb {
	const (
		StateDefault uint8 = iota
		StateUrl
	)

	users := make(map[int64]uint8)
	lock := sync.RWMutex{}

	return func(ctx *th.Context, update telego.Update) error {
		from := update.Message.From
		var text string

		lock.RLock()
		userState := users[from.ID]
		lock.RUnlock()

		switch userState {
		case StateDefault:
			text = "Enter article url:"
			lock.Lock()
			users[from.ID] = StateUrl
			lock.Unlock()

		case StateUrl:
			// url := update.Message.Text
			// articleRepo.Create(&article.Article{: url})
			// ctx.Bot().SendMessage(ctx, tu.Message(
			// 	tu.ID(update.Message.Chat.ID),
			// 	fmt.Sprintf("Hello %s!", update.Message.From.FirstName),
			// ))
			lock.Lock()
			delete(users, from.ID)
			lock.Unlock()
		}
		return nil
	}
}
