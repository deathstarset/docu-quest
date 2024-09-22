-- name: FindMessageByID :one
SELECT * FROM messages WHERE id = $1 LIMIT 1;

-- name: FindMessagesByConversationID :many
SELECT * FROM messages WHERE conversation_id = $1;

-- name: FindMessages :many
SELECT * FROM messages;

-- name: AddMessage :one
INSERT INTO messages (conversation_id, content, sender) VALUES ($1, $2, $3) RETURNING *;

-- name: RemoveMessage :exec
DELETE FROM messages WHERE id = $1;


