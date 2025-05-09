package user

import (
	"errors"
	"fmt"
	articleRepo "go-articles-manager-bot/internal/repositories/article"
	"time"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func (uh *UserHandler) NewGetUserStatHandler() th.Handler {
	return func(ctx *th.Context, update telego.Update) error {
		var text string
		now := time.Now()
		monthPeriod := []time.Time{now.AddDate(0, -1, 0), now}

		allCount, err := uh.articleRepo.GetArticlesCountByPeriod(update.Message.From.ID, true, monthPeriod[0], monthPeriod[1])
		notReadCount, err := uh.articleRepo.GetArticlesCountByPeriod(update.Message.From.ID, false, monthPeriod[0], monthPeriod[1])

		if err != nil {
			if errors.Is(err, articleRepo.ErrNotFound) {
				text = "Articles not found"
			}
			text = "Internal error"
		} else {
			text = fmt.Sprintf("In the last month:\n\nAricles added: %d\nArticles read: %d", allCount, allCount-notReadCount)
		}

		ctx.Bot().SendMessage(ctx, tu.Message(update.Message.Chat.ChatID(), text))
		return nil
	}
}
