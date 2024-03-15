package constants

type CreateDatabaseQuery string

const (
	ExpenseCalculatorCreateDatabaseQuery CreateDatabaseQuery = `CREATE TABLE IF NOT EXISTS 
    expense (
    id INTEGER PRIMARY KEY,
    amount DOUBLE PRECISION NOT NULL,
    created_at TEXT NOT NULL)`
)

func (c CreateDatabaseQuery) String() string {
	return string(c)
}
