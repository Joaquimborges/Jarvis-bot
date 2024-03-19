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

type FindExpensesByMonth struct {
	repository repository.ExpenseCalculator
	database   *sql.DB
}

func NewFindExpensesByMonthUsecase(database *sql.DB) *FindExpensesByMonth {
	return &FindExpensesByMonth{
		database:   database,
		repository: expense_calculator.NewExpenseCalculatorRepository(database),
	}
}

func (*FindExpensesByMonth) IsValid(message string) bool {
	return util.ContainsValue(
		message,
		[]string{
			"meses atrás",
			"meses atras",
			"mês atrás",
			"mes atras",
			" month expenses",
		},
	)
}

func (f *FindExpensesByMonth) BuildResponse(message, _ string) string {
	if f.database == nil {
		return fmt.Sprintf(
			constants.ImportForgotMessage,
			"database",
			"bot.WithDatabase()",
		)
	}

	qtt := util.GetNumberValueFromMessage(message)
	now, before := util.BuildComparableTime(0, -qtt)
	resp, err := f.repository.Select(constants.GetExpenseByDateNumber.String(), before, now)
	if err != nil {
		return fmt.Sprintf("[usecase.Select.all]Error was fount: %v", err)
	}
	logger.Usecase("FindExpensesByMonth")
	return common.BuildExpenseResponse(resp)
}
