-- name: GetAuthorByID :one
SELECT * FROM authors WHERE id = $1 LIMIT 1;

-- name: GetAuthorByEmail :one
SELECT * FROM authors WHERE email = $1 LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM authors ORDER BY name;

-- name: CreateAuthor :one
INSERT INTO authors (name, email) 
VALUES ($1, $2) 
RETURNING *;

-- name: UpdateAuthor :exec
UPDATE authors 
SET name = $2, email = $3, updated_at = NOW() 
WHERE id = $1;

-- name: DeleteAuthor :exec
DELETE FROM authors WHERE id = $1;
