-- +goose Up
CREATE EXTENSION IF NOT EXISTS vector;

CREATE TABLE
  IF NOT EXISTS embeddings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    content_id UUID NOT NULL,
    text TEXT NOT NULL,
    embedding VECTOR (768) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (content_id) REFERENCES extracted_contents (id) ON DELETE CASCADE
  );

CREATE INDEX ON embeddings USING hnsw (embedding vector_l2_ops);

-- +goose Down
DROP TABLE IF EXISTS embeddings;

DROP EXTENSION IF EXISTS vector;