package expense_calculator

import (
	"database/sql"
	"github.com/Joaquimborges/jarvis-bot/pkg/bot/logger"
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
	tx := c.initDatabase()
	stmt, err := tx.Prepare(constants.InsertExpense.String())
	if err != nil {
		return err
	}
	defer c.closeStatement(stmt, "Save()")

	_, err = stmt.Exec(data.Name, data.Amount, data.Description, data.Date)
	if err != nil {
		return err
	}

	er := tx.Commit()
	if er != nil {
		return er
	}
	return nil
}

func (c *calculator) Select(query string, args ...any) ([]*entities.ExpenseCalculatorBody, error) {
	tx := c.initDatabase()
	rows, err := tx.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer c.closeRows(rows, "Select()")
	decode := make([]*entities.ExpenseCalculatorBody, 0)
	var data entities.ExpenseCalculatorBody

	for rows.Next() {
		if er := rows.Scan(
			&data.Name,
			&data.Amount,
			&data.Description,
			&data.Date,
		); er != nil {
			return nil, er
		}
		decode = append(decode, &entities.ExpenseCalculatorBody{
			Description: data.Description,
			Name:        data.Name,
			Amount:      data.Amount,
			Date:        data.Date,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return decode, nil
}

func (c *calculator) initDatabase() *sql.Tx {
	tx, err := c.db.Begin()
	if err != nil {
		logger.Warn(
			"expense_calculator.closeDatabase(): %v",
			err.Error(),
		)
		return nil
	}
	return tx
}

func (*calculator) closeStatement(stmt *sql.Stmt, fn string) {
	err := stmt.Close()
	if err != nil {
		logger.Warn(
			"expense_calculator_repository.%s.closeStatement(): %v",
			fn,
			err.Error(),
		)
	}
}

func (*calculator) closeRows(rows *sql.Rows, fn string) {
	err := rows.Close()
	if err != nil {
		logger.Warn(
			"expense_calculator_repository.%s.closeRows(): %v",
			fn,
			err.Error(),
		)
	}
}
