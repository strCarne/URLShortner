-- name: CreateLink :one
INSERT INTO links (id)
VALUES ($1)
RETURNING *;