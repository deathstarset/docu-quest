-- +goose Up
CREATE TYPE user_type AS ENUM ('admin', 'user');

ALTER TABLE users
ADD COLUMN role user_type NOT NULL DEFAULT 'user';

-- +goose Down
ALTER TABLE users
DROP COLUMN role;

DROP TYPE user_type;