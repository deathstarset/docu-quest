-- name: FindConversationByID :one
SELECT * FROM conversations WHERE id = $1 LIMIT 1;

-- name: FindConversationsByUserID :many
SELECT * FROM conversations WHERE user_id = $1;

-- name: FindConversations :many
SELECT * FROM conversations ORDER BY started_at;

-- name: AddConversation :one
INSERT INTO conversations (user_id, document_id) VALUES ($1, $2) RETURNING *;

-- name: RemoveConversation :exec
DELETE FROM conversations WHERE id = $1;

-- name: FindUserConversation :one
SELECT * FROM conversations WHERE id = $1 AND user_id = $2 LIMIT 1;

-- name: RemoveUserConversation :exec
DELETE FROM conversations WHERE id = $1 AND user_id = $2;


