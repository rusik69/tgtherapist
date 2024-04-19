package tgtherapist

import (
	"github.com/rusik69/tgtherapist/pkg/telegram"
	"github.com/rusik69/tgtherapist/pkg/types"
	openai "github.com/sashabaranov/go-openai"
)

func main() {
	types.OpenAIClient = openai.NewClient(types.OpenAIToken)
	telegram.Serve()
}

func init() {
	types.ParseEnv()
}
