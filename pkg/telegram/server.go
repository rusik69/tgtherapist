package telegram

import (
	"context"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/rusik69/tgtherapist/pkg/types"
)

// Serve serves the Telegram bot.
func Serve() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}
	b, err := bot.New(types.TelegramToken, opts...)
	if err != nil {
		panic(err)
	}
	b.Start(ctx)
}
