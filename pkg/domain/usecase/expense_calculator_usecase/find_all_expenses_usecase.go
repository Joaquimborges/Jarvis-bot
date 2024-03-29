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

type FindAllExpenses struct {
	repository repository.ExpenseCalculator
	database   *sql.DB
}

func NewFindAllExpenseUsecase(database *sql.DB) *FindAllExpenses {
	return &FindAllExpenses{
		database:   database,
		repository: expense_calculator.NewExpenseCalculatorRepository(database),
	}
}

func (*FindAllExpenses) IsValid(message string) bool {
	return util.ContainsValue(
		message,
		[]string{
			"ver gastos",
			"todos os gastos",
			"expenses",
		},
	)
}

func (f *FindAllExpenses) BuildResponse(_, _ string) string {
	if f.database == nil {
		return fmt.Sprintf(
			constants.ImportForgotMessage,
			"database",
			"bot.WithDatabase()",
		)
	}

	resp, err := f.repository.Select(constants.GetAllExpense.String())
	if err != nil {
		return fmt.Sprintf("[usecase.Select.all]Error was fount: %v", err)
	}
	logger.Usecase("FindAllExpenses")
	return common.BuildExpenseResponse(resp)
}
