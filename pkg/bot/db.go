package bot

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"path"
	"runtime"
)

func InitDatabase(dbName string) (*sql.DB, error) {
	dir, er := getRootDir()
	if er != nil {
		return nil, er
	}

	return sql.Open(
		"sqlite3",
		fmt.Sprintf("%s/%s", dir, dbName))
}

func getRootDir() (string, error) {
	_, filename, _, found := runtime.Caller(0)
	if !found {
		return "", errors.New("unable to read current file name")
	}
	pwd := path.Dir(filename)
	return path.Join(pwd, "..", ".."), nil
}
