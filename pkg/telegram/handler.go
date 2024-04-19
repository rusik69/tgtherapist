package telegram

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/rusik69/tgtherapist/pkg/openai"
)

// handler is the default handler for the bot.
func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	input := update.Message.Text
	output, err := openai.Generate(input)
	if err != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   err.Error(),
		})
		return
	}
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   output,
	})
}
