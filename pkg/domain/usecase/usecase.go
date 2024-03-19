package usecase

import (
	"database/sql"
	"github.com/Joaquimborges/jarvis-bot/pkg/domain/usecase/chat_gpt_usecase"
	"github.com/Joaquimborges/jarvis-bot/pkg/domain/usecase/exchange"
	"github.com/Joaquimborges/jarvis-bot/pkg/domain/usecase/expense_calculator_usecase"
	"github.com/Joaquimborges/jarvis-bot/pkg/domain/usecase/wake_up_server"
	"github.com/Joaquimborges/jarvis-bot/pkg/gateway/open_ai"
	"github.com/Joaquimborges/jarvis-bot/pkg/gateway/rest"
	"github.com/Joaquimborges/jarvis-bot/pkg/util"
)

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
	genericUsecaseSlice []Usecase
	expensesBySlice     []Usecase
}

func NewJarvisUsecase(
	gpt open_ai.OpenAI,
	database *sql.DB,
	client rest.Waitress,
	testServerURLs ...string,
) UCBuilder {

	return &jarvisUsecase{
		genericUsecaseSlice: []Usecase{
			chat_gpt_usecase.NewAskOpenAI(gpt),
			exchange.NewExchangeUsecase(client),
			wake_up_server.NewWakeServersUsecase(client, testServerURLs...),
		},
		expensesBySlice: []Usecase{
			expense_calculator_usecase.NewSaveExpenseUsecase(database),
			expense_calculator_usecase.NewFindAllExpenseUsecase(database),
			expense_calculator_usecase.NewFindExpensesByMonthUsecase(database),
			expense_calculator_usecase.NewFindExpensesByDaysAgoUsecase(database),
		},
	}
}

func (uc *jarvisUsecase) BuildResponseContext(message, sender string) string {
	compilerSlice := util.ExpenseRegexCompiler.FindAllString(message, -1)
	if len(compilerSlice) < 1 {
		return uc.buildResponse(message, sender, uc.genericUsecaseSlice)
	}
	return uc.buildResponse(message, sender, uc.expensesBySlice)
}

func (*jarvisUsecase) buildResponse(message, sender string, cases []Usecase) string {
	for _, usecase := range cases {
		if usecase.IsValid(message) {
			return usecase.BuildResponse(message, sender)
		}
	}
	return notFoundContextKey
}
