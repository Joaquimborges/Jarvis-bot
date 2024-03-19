package common

import (
	"fmt"
	"github.com/Joaquimborges/jarvis-bot/pkg/domain/entities"
	"github.com/Joaquimborges/jarvis-bot/pkg/util"
)

func BuildExpenseResponse(data []*entities.ExpenseCalculatorBody) string {
	result := ""
	for _, expense := range data {
		result += fmt.Sprintf(
			"Description: %s\nAmount: R$%.2f\nDate: %s\nFrom: %s\n#-------------------------#\n\n",
			expense.Description,
			expense.Amount,
			util.ParseDate(expense.Date),
			expense.Name,
		)
	}
	return result
}
