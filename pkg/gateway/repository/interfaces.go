package repository

import "github.com/Joaquimborges/jarvis-bot/pkg/domain/entities"

type ExpenseCalculator interface {
	Select(query string, id string) (*entities.ExpenseCalculatorBody, error)
	Save(data *entities.ExpenseCalculatorBody) error
}
