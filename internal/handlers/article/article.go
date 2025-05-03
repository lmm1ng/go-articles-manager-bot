package article

import (
	"errors"
	"go-articles-manager-bot/internal/handlers"
	"go-articles-manager-bot/internal/middlewares"
	"go-articles-manager-bot/internal/models"
	"net/http"
	"time"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"golang.org/x/net/html"
)

type ArticleRepository interface {
	Create(*models.Article) error
}

type UserRepository interface {
	GetByTgUsername(string) (*models.User, error)
}

//	func NewRandomArticleHandler(articleRepo Article) handlers.Cb {
//		return func(ctx *th.Context, update telego.Update) error {
//			ctx.Bot().SendMessage(ctx, tu.Message(
//				tu.ID(update.Message.Chat.ID),
//				fmt.Sprintf("Hello %s!", update.Message.From.FirstName),
//			))
//			return nil
//		}
//	}

func NewEnterCreateArticleHandler() handlers.Cb {
	return func(ctx *th.Context, update telego.Update) error {
		scenesManager := ctx.Value(middlewares.ScenesManagerKey).(*models.ScenesManager)

		scenesManager.Mutex.RLock()
		curScene := scenesManager.Users[update.Message.From.ID]
		scenesManager.Mutex.RUnlock()

		if curScene != models.NoScene {
			return ctx.Next(update)
		}

		scenesManager.Mutex.Lock()
		scenesManager.Users[update.Message.From.ID] = models.StateAddUrl
		scenesManager.Mutex.Unlock()

		ctx.Bot().SendMessage(ctx, tu.Message(update.Message.Chat.ChatID(), "Enter article url:"))
		return nil
	}
}

func NewCreateArticleHandler(articleRepo ArticleRepository, userRepo UserRepository) handlers.Cb {
	client := http.Client{
		Timeout: time.Second,
	}

	return func(ctx *th.Context, update telego.Update) error {
		from := update.Message.From
		var text string

		scenesManager := ctx.Value("scenesManager").(*models.ScenesManager)

		scenesManager.Mutex.RLock()
		userState := scenesManager.Users[from.ID]
		scenesManager.Mutex.RUnlock()

		if userState == models.NoScene {
			return ctx.Next(update)
		}

		defer func() {
			ctx.Bot().SendMessage(ctx, tu.Message(update.Message.Chat.ChatID(), text))
		}()

		user, err := userRepo.GetByTgUsername(from.Username)

		if err != nil {
			text = "User not found"
			return nil
		}

		url := update.Message.Text
		resp, err := client.Get(url)

		if err != nil {
			text = "Url not valid"
			return nil
		}

		defer func() {
			resp.Body.Close()
		}()

		title, err := getTitle(resp)

		if err != nil {
			text = "Url not valid (no title)"
			return nil
		}

		err = articleRepo.Create(&models.Article{Url: url, Title: title, UserId: user.Id})

		if err != nil {
			text = "Article not created"
			return nil
		}

		text = "Article created successfully"

		scenesManager.Mutex.Lock()
		delete(scenesManager.Users, from.ID)
		scenesManager.Mutex.Unlock()

		return nil
	}
}

func getTitle(resp *http.Response) (string, error) {
	t := html.NewTokenizer(resp.Body)
	for {
		cur := t.Next()
		if cur == html.ErrorToken {
			return "", errors.New("Title not found")
		}

		if cur != html.StartTagToken {
			continue
		}
		token := t.Token()
		if token.Data != "meta" {
			continue
		}

		var ok bool

		for _, attr := range token.Attr {
			if attr.Key == "property" && attr.Val == "og:title" {
				ok = true
			}

			if attr.Key == "content" && ok {
				return attr.Val, nil
			}
		}
	}
}
