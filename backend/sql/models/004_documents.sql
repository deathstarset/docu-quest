-- +goose Up
CREATE TABLE
  IF NOT EXISTS documents (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    file_path VARCHAR(255) NOT NULL,
    uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    uploaded_by UUID NOT NULL,
    FOREIGN KEY (uploaded_by) REFERENCES users (id) ON DELETE CASCADE
  );

-- +goose Down
DROP TABLE documents;