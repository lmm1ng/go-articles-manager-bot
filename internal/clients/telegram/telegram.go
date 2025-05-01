package telegram

import (
	"context"
	"log/slog"

	"go-articles-manager-bot/internal/handlers"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

type TelegramClient struct {
	bot     *telego.Bot
	updates <-chan telego.Update
}

func New(token string, logger *slog.Logger) (*TelegramClient, error) {
	bot, err := telego.NewBot(token)
	ctx := context.Background()

	defer logger.Info("Telegram bot instance created successfully")

	if err != nil {
		logger.Error("Failed to create Telegram bot", "error", err)
		return nil, err
	}

	updates, _ := bot.UpdatesViaLongPolling(ctx, nil)

	return &TelegramClient{bot: bot, updates: updates}, nil
}

func (c *TelegramClient) RunHandlers(handlers []handlers.Handler) error {
	bh, _ := th.NewBotHandler(c.bot, c.updates)

	defer func() {
		bh.Stop()
	}()

	for _, h := range handlers {
		bh.Handle(h.Cb, h.Predicate)
	}

	bh.Start()

	return nil
}
