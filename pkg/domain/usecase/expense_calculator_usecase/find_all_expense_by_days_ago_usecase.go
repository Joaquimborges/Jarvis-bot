package expense_calculator_usecase

import (
	"database/sql"
	"fmt"
	"github.com/Joaquimborges/jarvis-bot/pkg/bot/logger"
	"github.com/Joaquimborges/jarvis-bot/pkg/domain/constants"
	"github.com/Joaquimborges/jarvis-bot/pkg/domain/usecase/expense_calculator_usecase/common"
	"github.com/Joaquimborges/jarvis-bot/pkg/gateway/repository"
	"github.com/Joaquimborges/jarvis-bot/pkg/gateway/repository/expense_calculator"
	"github.com/Joaquimborges/jarvis-bot/pkg/util"
)

type FindExpensesByDaysAgo struct {
	repository repository.ExpenseCalculator
	database   *sql.DB
}

func NewFindExpensesByDaysAgoUsecase(database *sql.DB) *FindExpensesByDaysAgo {
	return &FindExpensesByDaysAgo{
		database:   database,
		repository: expense_calculator.NewExpenseCalculatorRepository(database),
	}
}

func (*FindExpensesByDaysAgo) IsValid(message string) bool {
	return util.ContainsValue(
		message,
		[]string{
			"dias atrás",
			"dias atras",
			"dia atrás",
			"dia atras",
			" days ago expenses",
		},
	)
}

func (f *FindExpensesByDaysAgo) BuildResponse(message, _ string) string {
	if f.database == nil {
		return fmt.Sprintf(
			constants.ImportForgotMessage,
			"database",
			"bot.WithDatabase()",
		)
	}

	qtt := util.GetNumberValueFromMessage(message)
	now, before := util.BuildComparableTime(-qtt, 0)
	resp, err := f.repository.Select(constants.GetExpenseByDateNumber.String(), before, now)
	if err != nil {
		return fmt.Sprintf("[usecase.Select.all]Error was fount: %v", err)
	}
	logger.Usecase("FindExpensesByDaysAgo")
	return common.BuildExpenseResponse(resp)
}
