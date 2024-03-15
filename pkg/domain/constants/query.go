package constants

type JarvisDbQuery string

func (j JarvisDbQuery) String() string {
	return string(j)
}

const (
	InsertExpense JarvisDbQuery = `INSERT INTO 
    			expense (amount, created_at)
				VALUES (?, ?)`
)
