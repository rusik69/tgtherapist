package openai

import (
	"context"

	"github.com/rusik69/tgtherapist/pkg/types"
	goopenai "github.com/sashabaranov/go-openai"
)

// Generate generates a text based on the prompt.
func Generate(prompt string) (string, error) {
	resp, err := types.OpenAIClient.CreateChatCompletion(
		context.Background(),
		goopenai.ChatCompletionRequest{
			Model: goopenai.GPT3Dot5Turbo,
			Messages: []goopenai.ChatCompletionMessage{
				{
					Role:    goopenai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
