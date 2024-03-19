package constants

type JarvisDbQuery string

func (j JarvisDbQuery) String() string {
	return string(j)
}

const (
	InsertExpense JarvisDbQuery = `INSERT INTO 
    			expense (name, amount, description, created_at)
				VALUES (?, ?, ?, ?)`

	GetAllExpense            JarvisDbQuery = `SELECT name, amount, description, created_at FROM expense ORDER BY amount DESC`
	GetExpenseByDateNumber   JarvisDbQuery = `SELECT name, amount, description, created_at FROM expense WHERE created_at BETWEEN ? AND ? ORDER BY amount DESC`
	GetExpenseByDateAndOwner JarvisDbQuery = `SELECT name, amount, description, created_at FROM expense WHERE created_at BETWEEN ? AND ? AND name = ? ORDER BY amount DESC`
)
