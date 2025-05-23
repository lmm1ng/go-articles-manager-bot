package article

import (
	"errors"
	"go-articles-manager-bot/internal/entities"
	userRepo "go-articles-manager-bot/internal/repositories/user"
	"net/http"
	"slices"
	"time"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"golang.org/x/net/html"
)

func (ah *ArticleHandler) NewEnterCreateArticleHandler() th.Handler {
	return func(ctx *th.Context, update telego.Update) error {
		ctx.Bot().SendMessage(ctx, tu.Message(update.Message.Chat.ChatID(), "Enter article url:"))
		return nil
	}
}

func (ah *ArticleHandler) NewCreateArticleHandler() th.Handler {
	return func(ctx *th.Context, update telego.Update) error {
		client := http.Client{
			Timeout: time.Second,
		}

		from := update.Message.From
		var text string

		defer func() {
			ctx.Bot().SendMessage(ctx, tu.Message(update.Message.Chat.ChatID(), text))
		}()

		u, err := ah.userRepo.GetByTgId(from.ID)

		if err != nil {
			if errors.Is(err, userRepo.ErrNotFound) {
				text = "User not found"
			}

			text = "Internal error"

			return err
		}

		url := update.Message.Text
		resp, err := client.Get(url)

		if err != nil {
			text = "Url not valid"
			return err
		}

		defer func() {
			resp.Body.Close()
		}()

		var title *string
		extractedTitle, _ := getTitle(resp)
		if extractedTitle != "" {
			title = &extractedTitle
		}

		err = ah.articleRepo.
			Create(
				&entities.Article{
					Url:    url,
					Title:  title,
					UserId: u.Id,
				},
			)

		if err != nil {
			text = "Article not created"
			return err
		}

		text = "Article created successfully"

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
			if slices.Contains([]string{"property", "name"}, attr.Key) && attr.Val == "og:title" {
				ok = true
			}

			if attr.Key == "content" && ok {
				return attr.Val, nil
			}
		}
	}
}
