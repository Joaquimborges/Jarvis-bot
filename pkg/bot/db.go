package bot

import (
	"database/sql"
)

func InitDatabase() (*sql.DB, error) {
	return sql.Open(
		"sqlite3",
		"../../jarvis.db",
	)
}
