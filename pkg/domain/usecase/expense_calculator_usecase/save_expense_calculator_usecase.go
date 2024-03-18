package expense_calculator_usecase

import (
	"database/sql"
	"fmt"
	"github.com/Joaquimborges/jarvis-bot/pkg/bot/logger"
	"github.com/Joaquimborges/jarvis-bot/pkg/domain/entities"
	"github.com/Joaquimborges/jarvis-bot/pkg/gateway/repository"
	"github.com/Joaquimborges/jarvis-bot/pkg/gateway/repository/expense_calculator"
	"github.com/Joaquimborges/jarvis-bot/pkg/util"
	"strings"
)

type SaveExpense struct {
	repository repository.ExpenseCalculator
	database   *sql.DB
}

func NewSaveExpenseUsecase(database *sql.DB) *SaveExpense {
	return &SaveExpense{
		database:   database,
		repository: expense_calculator.NewExpenseCalculatorRepository(database),
	}
}

func (*SaveExpense) IsValid(message string) bool {
	return util.ContainsValue(
		message,
		[]string{
			"gastei",
			"comprei",
			"anota nos gastos",
			"gastos externos",
			"acabei de gastar",
		},
	)
}

func (s *SaveExpense) BuildResponse(message, sender string) string {
	if s.database == nil {
		return fmt.Sprintf(
			"You forgot to import database dependency, \nuse the %s option",
			"bot.WithDatabase()",
		)
	}

	data := strings.Split(message, ", ")
	amount := data[1]
	description := data[2]
	payload, err := entities.NewExpenseCalculatorBody(sender, amount, description)
	if err != nil {
		return fmt.Sprintf("[usecase.NewExpenseCalculatorBody]Error was fount: %v", err)
	}

	if er := s.repository.Save(payload); er != nil {
		return fmt.Sprintf("[usecase.SaveExpense()]Error was fount: %v", err)
	}
	logger.Usecase("SaveExpense")
	return fmt.Sprintf("I just saved the amount: %s \nwith the name: %s \nin the external expenses list",
		amount,
		description,
	)
}
