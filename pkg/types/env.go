package types

import (
	"os"

	goopenai "github.com/sashabaranov/go-openai"
)

// Telegram and OpenAI tokens.
var TelegramToken, OpenAIToken string
var OpenAIClient *goopenai.Client

// ParseEnv parses the environment variables.
func ParseEnv() {
	TelegramToken = os.Getenv("TELEGRAM_TOKEN")
	if TelegramToken == "" {
		panic("TELEGRAM_TOKEN is required")
	}
	OpenAIToken = os.Getenv("OPENAI_TOKEN")
	if OpenAIToken == "" {
		panic("OPENAI_TOKEN is required")
	}
}
