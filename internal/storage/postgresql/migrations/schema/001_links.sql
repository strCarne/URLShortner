-- +goose Up

CREATE TABLE urls (
    id UUID NOT NULL PRIMARY KEY,
    alias TEXT NOT NULL UNIQUE,
    url TEXT NOT NULL
);
CREATE INDEX IF NOT EXISTS idx_alias on urls(alias);

-- +goose Down

DROP TABLE urls;