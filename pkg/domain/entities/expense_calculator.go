package entities

import (
	"github.com/Joaquimborges/jarvis-bot/pkg/util"
	"os"
	"strconv"
	"time"
)

type ExpenseCalculatorBody struct {
	Description string
	Name        string
	Amount      float64
	Date        time.Time
}

func NewExpenseCalculatorBody(name, amount, description string) (*ExpenseCalculatorBody, error) {
	floatAmount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return nil, err
	}

	return &ExpenseCalculatorBody{
			Name:   name,
			Amount: floatAmount,
			Date: util.CreateLocalTime(
				os.Getenv("TIME_LOCATION"),
			),
			Description: description,
		},
		nil
}
