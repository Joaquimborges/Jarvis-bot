package constants

type CreateDatabaseQuery string

const (
	ExpenseCalculatorCreateDatabaseQuery CreateDatabaseQuery = `CREATE TABLE IF NOT EXISTS 
    expense (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    amount DOUBLE PRECISION NOT NULL,
    created_at TEXT NOT NULL)`
)
