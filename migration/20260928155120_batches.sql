-- +goose Up
CREATE TABLE IF NOT EXISTS batches
(
    user_id       BIGINT      NOT NULL,
    collection_id INTEGER     NOT NULL,
    batch_id      UUID        NOT NULL,
    expiry        TIMESTAMPTZ NOT NULL,
    PRIMARY KEY (user_id, collection_id, batch_id),
    FOREIGN KEY (user_id, collection_id) REFERENCES user_collections (user_id, collection_id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS batches;
