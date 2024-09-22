-- name: FindConversationByID :one
SELECT * FROM conversations WHERE id = $1 LIMIT 1;

-- name: FindConversationsByUserID :many
SELECT * FROM conversations WHERE user_id = $1;

-- name: FindConversations :many
SELECT * FROM conversations ORDER BY started_at;

-- name: AddConversation :one
INSERT INTO conversations (user_id) VALUES ($1) RETURNING *;

-- name: RemoveConversation :exec
DELETE FROM conversations WHERE id = $1;

