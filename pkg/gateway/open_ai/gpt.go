package open_ai

import (
	"context"
	"errors"
	"github.com/sashabaranov/go-openai"
	"io"
)

//go:generate mockgen -source ./gpt.go -destination ../../internal/mocks/gateway/openai_mock.go -package mocks_gateway
type OpenAI interface {
	GetMessageContext(ctx context.Context, query string) (string, error)
}

type openIA struct {
	client      *openai.Client
	openaiModel string
	maxTokens   int
}

func NewOpenIAClient(openAIModel, openaiKey string, maxTokens int) OpenAI {
	return &openIA{
		client: openai.NewClient(
			openaiKey,
		),
		maxTokens:   maxTokens,
		openaiModel: openAIModel,
	}
}

func (o *openIA) GetMessageContext(ctx context.Context, query string) (string, error) {
	message := o.buildNewUserMessage(query)
	request := o.buildNewChatCompletionRequest(message)

	stream, err := o.client.CreateChatCompletionStream(ctx, request)
	if err != nil {
		return "", err
	}
	defer stream.Close()

	openIaContext := ""
	for {
		response, er := stream.Recv()
		if er != nil {
			if errors.Is(er, io.EOF) {
				return openIaContext, nil
			}
			return "", er
		}
		openIaContext += response.Choices[0].Delta.Content
	}
}

func (*openIA) buildNewUserMessage(content string) []openai.ChatCompletionMessage {
	return []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: content,
		},
	}
}

func (o *openIA) buildNewChatCompletionRequest(message []openai.ChatCompletionMessage) openai.ChatCompletionRequest {
	return openai.ChatCompletionRequest{
		Model:     o.openaiModel,
		MaxTokens: o.maxTokens,
		Messages:  message,
		Stream:    true,
	}
}
