package postgresql

import (
	"database/sql"
	"os"

	"github.com/strCarne/URLShortner/internal/storage/postgresql/sqlc/generated"
)

type Storage struct {
	DB *generated.Queries
}

func New() (*Storage, error) {
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil { 
		return nil, err
	}
	return &Storage{generated.New(db)}, nil
}