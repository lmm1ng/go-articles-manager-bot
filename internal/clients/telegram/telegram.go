package telegram

import (
	"context"
	"log/slog"

	"go-articles-manager-bot/internal/handlers"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

type telegramClient struct {
	bot     *telego.Bot
	updates <-chan telego.Update
}

func MustNew(token string, logger *slog.Logger) *telegramClient {
	bot, err := telego.NewBot(token)
	ctx := context.Background()

	if err != nil {
		panic("Failed to create Telegram bot")
	}

	defer logger.Info("Telegram bot instance created successfully")

	updates, _ := bot.UpdatesViaLongPolling(ctx, nil)

	return &telegramClient{bot: bot, updates: updates}
}

func (c *telegramClient) RunHandlers(handlers []handlers.Handler, middlewares []handlers.Cb) error {
	bh, _ := th.NewBotHandler(c.bot, c.updates)

	defer bh.Stop()

	for _, m := range middlewares {
		bh.Use(m)
	}

	for _, h := range handlers {
		bh.Handle(h.Cb, h.Predicate)
	}

	bh.Start()

	return nil
}
