-- +goose Up
CREATE TABLE
  IF NOT EXISTS conversations (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    started_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    user_id UUID NOT NULL,
    ended_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
  );

-- +goose Down
DROP TABLE conversations;