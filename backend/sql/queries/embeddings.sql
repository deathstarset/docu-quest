-- name: FindEmbeddingByID :one
SELECT * FROM embeddings WHERE id = $1 LIMIT 1;

-- name: FindEmbeddingsByContentID :many
SELECT * FROM embeddings WHERE content_id = $1;

-- name: AddEmbedding :one
INSERT INTO embeddings (content_id, text, embedding) VALUES ($1, $2, $3) RETURNING *;

-- name: RemoveEmbedding :exec
DELETE FROM embeddings WHERE id = $1;

-- name: FindSimilarVec :many
SELECT * FROM embeddings JOIN extracted_contents ON embeddings.content_id = extracted_contents.id WHERE extracted_contents.document_id = $1 ORDER BY embedding <-> $2 LIMIT 10;

