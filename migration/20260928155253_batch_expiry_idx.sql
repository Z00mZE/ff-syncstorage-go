-- +goose Up
CREATE INDEX IF NOT EXISTS batch_expiry_idx ON batches (user_id, collection_id, expiry);

-- +goose Down
DROP INDEX IF EXISTS batch_expiry_idx;
