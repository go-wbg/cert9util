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

func (db *SQLiteDB) Close() {
	db.DB.Close()
}

func (db *SQLiteDB) Columns(table string) ([]string, error) {
	rows, err := db.DB.Query("SELECT * FROM " + table + ";")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	return cols, nil
}
