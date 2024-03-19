package common

import (
	"fmt"
	"github.com/Joaquimborges/jarvis-bot/pkg/domain/entities"
	"github.com/Joaquimborges/jarvis-bot/pkg/util"
)

func BuildExpenseResponse(data []*entities.ExpenseCalculatorBody) string {
	var (
		total  float64
		result string
	)
	for _, expense := range data {
		result += fmt.Sprintf(
			"Description: %s\nAmount: R$%.2f\nDate: %s\nFrom: %s\n#-------------------------#\n\n",
			expense.Description,
			expense.Amount,
			util.ParseDate(expense.Date),
			expense.Name,
		)
		total += expense.Amount
	}
	result += fmt.Sprintf("#*------------*#\nTotal: R$%.2f\n#*----------*#", total)
	return result
}
