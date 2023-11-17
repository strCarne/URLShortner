package postgresql

import (
	"bytes"
	"database/sql"
	"os"
	"os/exec"
	"strings"

	_ "github.com/lib/pq"
	"github.com/strCarne/URLShortner/internal/storage/postgresql/sqlc/generated"
	"github.com/strCarne/URLShortner/pkg/wraperr"
)

type Storage struct {
	DB *generated.Queries
}

func New() (*Storage, error) {
	const op = "storage.postgresql.New"

	cmd := exec.Cmd{Path: "python3", Args: []string{"goose.py", "--up"}}
	out, err := cmd.Output()
	if err != nil {
		return nil, wraperr.Wrap(op, err)
	}
	if bytes.Contains(out, []byte("ERROR")) {
		return nil, wraperr.Wrap(strings.Join([]string{op, string(out)}, ": "), err)
	}

	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, wraperr.Wrap(op, err)
	}
	return &Storage{generated.New(db)}, nil
}
