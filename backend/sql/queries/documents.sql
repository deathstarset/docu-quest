-- name: FindDocumentByID :one
SELECT * FROM documents WHERE id = $1 LIMIT 1;

-- name: FindDocumentsByUserID :many
SELECT * FROM documents WHERE uploaded_by = $1;

-- name: FindDocuments :many
SELECT * FROM documents;

-- name: AddDocument :one
INSERT INTO documents (file_path, uploaded_by) VALUES ($1, $2) RETURNING *;

-- name: RemoveDocument :exec
DELETE FROM documents WHERE id = $1;