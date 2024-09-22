-- name: FindExtractedContentByID :one
SELECT * FROM extracted_contents WHERE id = $1 LIMIT 1;

-- name: FindExtractedContents :many
SELECT * FROM extracted_contents ORDER BY created_at;

-- name: AddExtractedContent :one
INSERT INTO extracted_contents (document_id, content) VALUES ($1, $2) RETURNING *;

-- name: RemoveExtractedContent :exec
DELETE FROM extracted_contents WHERE id = $1;

