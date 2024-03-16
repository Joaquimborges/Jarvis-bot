package repository

import "github.com/Joaquimborges/jarvis-bot/pkg/domain/entities"

type ExpenseCalculator interface {
	Select(query string) ([]*entities.ExpenseCalculatorBody, error)
	Save(data *entities.ExpenseCalculatorBody) error
}
