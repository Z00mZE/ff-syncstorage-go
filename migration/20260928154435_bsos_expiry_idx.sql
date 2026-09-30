-- +goose Up
CREATE INDEX IF NOT EXISTS bsos_expiry_idx ON bsos (user_id, collection_id, expiry);

-- +goose Down
DROP INDEX IF EXISTS bsos_expiry_idx;
