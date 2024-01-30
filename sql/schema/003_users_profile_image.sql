-- +goose Up
ALTER TABLE users
ADD COLUMN profile_image BYTEA DEFAULT NULL;

-- +goose Down
ALTER TABLE users
DROP COLUMN profile_image;