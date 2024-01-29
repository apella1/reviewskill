-- +goose Up
ALTER TABLE users
ALTER COLUMN email TYPE VARCHAR(254);

-- +goose Down
ALTER TABLE users
ALTER COLUMN email TYPE TEXT;