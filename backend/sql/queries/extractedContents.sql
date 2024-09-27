-- name: FindExtractedContentByID :one
SELECT * FROM extracted_contents WHERE id = $1 LIMIT 1;

-- name: FindExtractedContents :many
SELECT * FROM extracted_contents ORDER BY created_at;

-- name: AddExtractedContent :one
INSERT INTO extracted_contents (document_id, content) VALUES ($1, $2) RETURNING *;

-- name: RemoveExtractedContent :exec
DELETE FROM extracted_contents WHERE id = $1;

-- name: RemoveUserExtractedContent :exec
DELETE FROM extracted_contents USING documents WHERE extracted_contents.document_id = documents.id AND documents.uploaded_by = $1 AND extracted_contents.id = $2;

-- name: FindUserExtractedContentByID :one
SELECT * FROM extracted_contents JOIN documents ON extracted_contents.document_id = documents.id WHERE documents.uploaded_by = $1 AND extracted_contents.id = $2;






