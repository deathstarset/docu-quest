-- +goose Up
CREATE TYPE sender_type AS ENUM ('bot', 'user');

CREATE TABLE
  IF NOT EXISTS messages (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    conversation_id UUID NOT NULL,
    content TEXT NOT NULL,
    sender sender_type NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CHECK (sender IN ('bot', 'user')),
    FOREIGN KEY (conversation_id) REFERENCES conversations (id) ON DELETE CASCADE
  );

-- +goose Down
ALTER TABLE messages
DROP COLUMN sender;

DROP TYPE sender_type CASCADE;

DROP TABLE messages;