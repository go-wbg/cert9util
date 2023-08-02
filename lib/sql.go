package cert9util

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type SQLiteDB struct {
	*sql.DB
}

func NewSQLiteDB(path string) (*SQLiteDB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	return &SQLiteDB{db}, nil
}
