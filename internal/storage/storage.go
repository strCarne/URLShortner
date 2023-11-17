package storage

import "errors"

var (
	ErrURLNotFound = errors.New("url not found")
	ErrURLExists   = errors.New("url exists")
	ErrPrepare     = errors.New("couldn't prepare the statement")
	ErrExec        = errors.New("couldn't exec the statement")
	ErrOpen        = errors.New("couldn't connect to the database")
	ErrQuery       = errors.New("couldn't make the query")
	ErrColumns     = errors.New("couldn't get columns")
)

type Storage interface {
	SaveURL(string, string) error
	GetURL(string) (string, error)
}
