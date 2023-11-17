package sqlite

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/mattn/go-sqlite3"
	"github.com/strCarne/URLShortner/internal/storage"
	"github.com/strCarne/URLShortner/pkg/wraperr"
)

const (
	idColumn = iota
	aliasColumn
	urlColumn
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, wraperr.WrapOp(op, storage.ErrOpen, err)
	}

	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS urls (
		id UUID NOT NULL PRIMARY KEY,
		alias TEXT NOT NULL UNIQUE,
		url TEXT NOT NULL
	);
	CREATE INDEX IF NOT EXISTS idx_alias on urls(alias);
	`)
	if err != nil {
		return nil, wraperr.WrapOp(op, storage.ErrPrepare, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, wraperr.WrapOp(op, storage.ErrExec, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) SaveURL(urlToSave, alias string) error {
	const op = "storage.sqlite.SaveURL"

	id := uuid.New()
	stmt, err := s.db.Prepare("INSERT INTO urls(id, url, alias) VALUES(?, ?, ?)")
	if err != nil {
		return wraperr.WrapOp(op, storage.ErrPrepare, err)
	}

	_, err = stmt.Exec(id, urlToSave, alias)
	if err != nil {
		if sqlLiteErr, ok := err.(sqlite3.Error); ok && sqlLiteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return wraperr.WrapOp(op, storage.ErrExec, storage.ErrURLExists)
		}
		return wraperr.Wrap(op, err)
	}

	return nil
}

func (s *Storage) GetURL(alias string) (string, error) {
	const op = "storage.sqlite"

	stmt, err := s.db.Prepare("SELECT * FROM urls WHERE alias = ?")
	if err != nil {
		return "", wraperr.WrapOp(op, storage.ErrPrepare, err)
	}

	var strURL string
	var id uuid.UUID
	err = stmt.QueryRow(alias).Scan(&id, &alias, &strURL)
	if err != nil {
		return "", wraperr.Wrap(op, err)
	}

	return strURL, nil
}
