package entities

import (
	"github.com/Joaquimborges/jarvis-bot/pkg/util"
	"os"
	"strconv"
)

type ExpenseCalculatorBody struct {
	Amount float64
	Date   string
}

func NewExpenseCalculatorBody(amount string) (*ExpenseCalculatorBody, error) {
	floatAmount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return nil, err
	}

	return &ExpenseCalculatorBody{
			Amount: floatAmount,
			Date: util.CreateNewStringLocalDate(
				os.Getenv("TIME_LOCATION"),
			),
		},
		nil
}
