package open_ia

import (
	"context"
	"errors"
	"github.com/sashabaranov/go-openai"
	"io"
	"os"
)

type OpenAI interface {
	GetMessageContext(ctx context.Context, query string) (string, error)
}

type openIA struct {
	client *openai.Client
}

func NewOpenIAClient() OpenAI {
	return &openIA{
		client: openai.NewClient(
			os.Getenv("OPEN_IA_API_KEY"),
		),
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

func (*openIA) buildNewChatCompletionRequest(message []openai.ChatCompletionMessage) openai.ChatCompletionRequest {
	return openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 500,
		Messages:  message,
		Stream:    true,
	}
}
