package expense_calculator

import (
	"database/sql"
	"github.com/Joaquimborges/jarvis-bot/pkg/bot/logger"
	"github.com/Joaquimborges/jarvis-bot/pkg/domain/constants"
	"github.com/Joaquimborges/jarvis-bot/pkg/domain/entities"
	"github.com/Joaquimborges/jarvis-bot/pkg/gateway/repository"
	"time"
)

type calculator struct {
	db *sql.DB
}

func NewExpenseCalculatorRepository(db *sql.DB) repository.ExpenseCalculator {
	return &calculator{db: db}
}

func (c *calculator) Save(data *entities.ExpenseCalculatorBody) error {
	defer c.closeDatabase()
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

func (c *calculator) Select(query string) ([]*entities.ExpenseCalculatorBody, error) {
	defer c.closeDatabase()
	tx := c.initDatabase()
	rows, err := tx.Query(query)
	if err != nil {
		return nil, err
	}
	defer c.closeRows(rows, "Select()")
	decode := make([]*entities.ExpenseCalculatorBody, 0)
	var name, description, date string
	var amount float64

	for rows.Next() {
		if er := rows.Scan(&name, &amount, &description, &date); er != nil {
			return nil, er
		}
		decode = append(decode, &entities.ExpenseCalculatorBody{
			Description: description,
			Name:        name,
			Amount:      amount,
			Date:        c.parseTime(date),
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return decode, nil
}

func (*calculator) parseTime(date string) time.Time {
	dateTime, er := time.Parse(time.RFC3339, date)
	if er != nil {
		return time.Now()
	}
	return dateTime
}

func (c *calculator) closeDatabase() {
	err := c.db.Close()
	if err != nil {
		logger.Warn(
			"expense_calculator.closeDatabase(): %v",
			err.Error(),
		)
		return
	}
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
