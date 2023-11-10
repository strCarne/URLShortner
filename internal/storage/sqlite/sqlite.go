package sqlite

import (
	"database/sql"
	"fmt"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.Mew"

	db, err := sql.Open("sqlite3", "")
	if err != nil {
		return nil, fmt.Errorf("%s: %s", op, err)
	}

	return &Storage{db}, nil
}