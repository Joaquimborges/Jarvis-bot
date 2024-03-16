package entities

import (
	"github.com/Joaquimborges/jarvis-bot/pkg/util"
	"os"
	"strconv"
)

type ExpenseCalculatorBody struct {
	Name   string
	Amount float64
	Date   string
}

func NewExpenseCalculatorBody(name, amount string) (*ExpenseCalculatorBody, error) {
	floatAmount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return nil, err
	}

	return &ExpenseCalculatorBody{
			Name:   name,
			Amount: floatAmount,
			Date: util.CreateNewStringLocalDate(
				os.Getenv("TIME_LOCATION"),
			),
		},
		nil
}
