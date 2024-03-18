package usecase

import (
	"database/sql"
	"github.com/Joaquimborges/jarvis-bot/pkg/domain/usecase/chat_gpt_usecase"
	"github.com/Joaquimborges/jarvis-bot/pkg/domain/usecase/exchange"
	"github.com/Joaquimborges/jarvis-bot/pkg/domain/usecase/expense_calculator_usecase"
	"github.com/Joaquimborges/jarvis-bot/pkg/domain/usecase/wake_up_server"
	"github.com/Joaquimborges/jarvis-bot/pkg/gateway/open_ai"
	"github.com/Joaquimborges/jarvis-bot/pkg/gateway/rest"
)

type ContextKey string

const (
	notFoundContextKey = "If you want to talk, write something more complete and starting with /ask ..."
)

//go:generate mockgen -source ./usecase.go -destination ../../internal/mocks/usecase/usecase_mock.go -package mocks_usecase
type Usecase interface {
	IsValid(message string) bool
	BuildResponse(message, sender string) string
}

type UCBuilder interface {
	BuildResponseContext(message, sender string) string
}

type jarvisUsecase struct {
	validUsecaseList []Usecase
}

func NewJarvisUsecase(
	gpt open_ai.OpenAI,
	database *sql.DB,
	client rest.Waitress,
	testServerURLs ...string,
) UCBuilder {

	return &jarvisUsecase{
		validUsecaseList: []Usecase{
			chat_gpt_usecase.NewAskOpenAI(gpt),
			exchange.NewExchangeUsecase(client),
			wake_up_server.NewWakeServersUsecase(client, testServerURLs...),
			expense_calculator_usecase.NewSaveExpenseUsecase(database),
			expense_calculator_usecase.NewFindAllExpenseUsecase(database),
		},
	}
}

func (uc *jarvisUsecase) BuildResponseContext(message, sender string) string {
	for _, usecase := range uc.validUsecaseList {
		if usecase.IsValid(message) {
			return usecase.BuildResponse(message, sender)
		}
	}
	return notFoundContextKey
}
