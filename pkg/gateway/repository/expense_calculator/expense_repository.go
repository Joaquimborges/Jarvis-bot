package expense_calculator

import (
	"database/sql"
	"fmt"
	"github.com/Joaquimborges/jarvis-bot/pkg/bot/config"
	"github.com/Joaquimborges/jarvis-bot/pkg/domain/constants"
	"github.com/Joaquimborges/jarvis-bot/pkg/domain/entities"
	"github.com/Joaquimborges/jarvis-bot/pkg/gateway/repository"
)

type calculator struct {
	db *sql.DB
}

func NewExpenseCalculatorRepository(db *sql.DB) repository.ExpenseCalculator {
	return &calculator{db: db}
}

func (c *calculator) Save(data *entities.ExpenseCalculatorBody) error {
	defer c.closeDatabase()

	tx, er := c.db.Begin()
	if er != nil {
		return er
	}

	stmt, err := tx.Prepare(constants.InsertExpense.String())
	if err != nil {
		return err
	}
	defer c.closeStatement(stmt, "Save()")

	_, err = stmt.Exec(data.Amount, data.Date)
	if err != nil {
		return err
	}

	er = tx.Commit()
	if er != nil {
		return er
	}
	return nil
}

func (c *calculator) Select(query string, id string) (*entities.ExpenseCalculatorBody, error) {
	defer c.closeDatabase()

	stmt, err := c.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer c.closeStatement(stmt, "Select()")

	var decode entities.ExpenseCalculatorBody
	err = stmt.QueryRow(id).Scan(&decode)
	if err != nil {
		return nil, err
	}
	return &decode, nil
}

func (c *calculator) closeDatabase() {
	err := c.db.Close()
	if err != nil {
		config.Logger.Fatal(
			fmt.Sprintf(
				"expense_calculator.closeDatabase(): %v",
				err,
			))
	}
}

func (*calculator) closeStatement(stmt *sql.Stmt, fn string) {
	err := stmt.Close()
	if err != nil {
		config.Logger.Fatal(
			fmt.Sprintf(
				"expense_calculator_repository.%s.closeStatement(): %v",
				fn,
				err,
			))
	}
}