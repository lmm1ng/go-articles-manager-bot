package telegram

import (
	"context"
	"go-articles-manager-bot/internal/handlers"
	"go-articles-manager-bot/internal/pkg/scenebuilder"
	"log/slog"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

type telegramClient struct {
	bot     *telego.Bot
	updates <-chan telego.Update
}

func MustNew(token string, logger *slog.Logger) *telegramClient {
	bot, err := telego.NewBot(token)

	if err != nil {
		panic("Failed to create Telegram bot by provided token")
	}

	ctx := context.Background()

	defer logger.Info("Telegram bot instance created successfully")

	updates, _ := bot.UpdatesViaLongPolling(ctx, nil)

	bot.SetMyCommands(
		ctx,
		&telego.SetMyCommandsParams{
			Commands: []telego.BotCommand{
				{Command: "/start", Description: "Register and start"},
				{Command: "/menu", Description: "Show main menu"},
			},
		},
	)

	return &telegramClient{bot: bot, updates: updates}
}

func (c *telegramClient) Run(
	handlers []handlers.Handler,
	middlewares []th.Handler,
	scenes []scenebuilder.Scene,
) error {
	bh, _ := th.NewBotHandler(c.bot, c.updates)

	defer bh.Stop()

	for _, m := range middlewares {
		bh.Use(m)
	}

	for _, h := range handlers {
		bh.Handle(h.Cb, h.Predicate)
	}

	for _, s := range scenes {
		s.Register(bh)
	}

	bh.Start()

	return nil
}
