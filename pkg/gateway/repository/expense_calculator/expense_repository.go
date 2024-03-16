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
	tx := c.initDatabase()
	stmt, err := tx.Prepare(constants.InsertExpense.String())
	if err != nil {
		return err
	}
	defer c.closeStatement(stmt, "Save()")

	_, err = stmt.Exec(data.Name, data.Amount, data.Date)
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
	var name, date string
	var amount float64

	for rows.Next() {
		if er := rows.Scan(&name, &amount, &date); er != nil {
			return nil, er
		}
		decode = append(decode, &entities.ExpenseCalculatorBody{
			Name:   name,
			Amount: amount,
			Date:   date,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return decode, nil
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

func (c *calculator) initDatabase() *sql.Tx {
	tx, err := c.db.Begin()
	if err != nil {
		config.Logger.Panic(
			fmt.Sprintf(
				"expense_calculator.closeDatabase(): %v",
				err,
			))
		return nil
	}
	return tx
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

func (*calculator) closeRows(rows *sql.Rows, fn string) {
	err := rows.Close()
	if err != nil {
		config.Logger.Fatal(
			fmt.Sprintf(
				"expense_calculator_repository.%s.closeRows(): %v",
				fn,
				err,
			))
	}
}
