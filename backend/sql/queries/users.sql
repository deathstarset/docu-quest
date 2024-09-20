-- name: FindUserByID :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: FindUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users ORDER BY username;

-- name: AddUser :one
INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING *;

-- name: RemoveUser :exec
DELETE FROM users WHERE id = $1;

-- name: EditUser :one
UPDATE users SET username = $1, email = $2, password = $3 WHERE id = $4 RETURNING *;