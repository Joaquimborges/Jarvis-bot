package constants

type JarvisDbQuery string

func (j JarvisDbQuery) String() string {
	return string(j)
}

const (
	InsertExpense JarvisDbQuery = `INSERT INTO 
    			expense (name, amount, description, created_at)
				VALUES (?, ?, ?, ?)`

	GetAllExpense JarvisDbQuery = `SELECT name, amount, description, created_at FROM expense`
)
