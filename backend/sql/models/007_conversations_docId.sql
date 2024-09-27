-- +goose Up
ALTER TABLE conversations
ADD COLUMN document_id UUID NOT NULL,
ADD FOREIGN KEY (document_id) REFERENCES documents (id) ON DELETE CASCADE;

-- +goose Down
ALTER TABLE conversations
DROP CONSTRAINT IF EXISTS fk_document_id;

ALTER TABLE conversations
DROP COLUMN document_id;