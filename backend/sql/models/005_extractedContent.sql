-- +goose Up
CREATE TABLE
  IF NOT EXISTS extracted_contents (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    document_id UUID NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (document_id) REFERENCES documents (id) ON DELETE CASCADE
  );

-- +goose Down
DROP TABLE extracted_contents;