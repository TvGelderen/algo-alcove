-- name: CreateUser :exec
INSERT INTO users (id, email, username, role, password_hash, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, timezone('utc', NOW()), timezone('utc', NOW()));

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;
