package chat_gpt_usecase

import (
	"context"
	"fmt"
	"github.com/Joaquimborges/jarvis-bot/pkg/bot/logger"
	"github.com/Joaquimborges/jarvis-bot/pkg/gateway/open_ai"
	"github.com/Joaquimborges/jarvis-bot/pkg/util"
	"strings"
)

type AskChatGpt struct {
	gpt open_ai.OpenAI
}

func NewAskOpenAI(gpt open_ai.OpenAI) *AskChatGpt {
	return &AskChatGpt{
		gpt: gpt,
	}
}

func (*AskChatGpt) IsValid(message string) bool {
	return util.ContainsValue(message,
		[]string{"/ask", "chat gpt", "gpt", "ask for gpt"},
	)
}

func (a *AskChatGpt) BuildResponse(message, _ string) string {
	if a.gpt == nil {
		logger.Usecase("[AskChatGpt without import]")
		return fmt.Sprintf(
			"You forgot to import the openai dependency, use the %s option",
			"bot.WithOpenAiIntegration()",
		)
	}

	message = strings.TrimPrefix(message, "/ask ")
	gptContext, err := a.gpt.GetMessageContext(context.Background(), message)
	if err != nil {
		return fmt.Sprintf("Error mounting context: %v", err)
	}
	logger.Usecase("[AskChatGpt]")
	return gptContext
}
