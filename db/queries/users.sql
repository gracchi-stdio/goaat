-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByGithubID :one
SELECT * FROM users
WHERE github_id = $1 LIMIT 1;

-- name: UpsertUser :one
INSERT INTO users (
    github_id,
    email,
    name,
    avatar_url
) VALUES (
    $1, $2, $3, $4
)
ON CONFLICT (github_id) DO UPDATE
SET
    email = EXCLUDED.email,
    name = EXCLUDED.name,
    avatar_url = EXCLUDED.avatar_url,
    updated_at = NOW()
RETURNING *;
