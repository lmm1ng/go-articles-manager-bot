// package user

// import (
// 	"go-articles-manager-bot/internal/middlewares"
// 	"go-articles-manager-bot/internal/models"

// 	"github.com/mymmrac/telego"
// 	th "github.com/mymmrac/telego/telegohandler"
// 	tu "github.com/mymmrac/telego/telegoutil"
// )

// func NewEnterAddUserDescHandler() th.Handler {
// 	return func(ctx *th.Context, update telego.Update) error {
// 		scenesManager := ctx.Value(middlewares.ScenesManagerKey).(*models.ScenesManager)

// 		scenesManager.Mutex.RLock()
// 		curScene := scenesManager.Users[update.Message.From.ID]
// 		scenesManager.Mutex.RUnlock()

// 		if curScene != models.NoScene {
// 			return ctx.Next(update)
// 		}

// 		scenesManager.Mutex.Lock()
// 		scenesManager.Users[update.Message.From.ID] = models.StateAddArticleUrl
// 		scenesManager.Mutex.Unlock()

// 		ctx.Bot().SendMessage(ctx, tu.Message(update.Message.Chat.ChatID(), "Enter article url:"))
// 		return nil
// 	}
// }
