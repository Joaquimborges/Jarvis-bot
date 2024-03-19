package repository

import "github.com/Joaquimborges/jarvis-bot/pkg/domain/entities"

//go:generate mockgen -source interfaces.go -destination ../../internal/mocks/gateway/expense_repository_mock.go -package mocks_gateway
type ExpenseCalculator interface {
	Select(query string, args ...any) ([]*entities.ExpenseCalculatorBody, error)
	Save(data *entities.ExpenseCalculatorBody) error
}
