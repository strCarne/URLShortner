-- +goose Up

CREATE TABLE links (
    id UUID NOT NULL PRIMARY KEY
);

-- +goose Down

DROP TABLE links;