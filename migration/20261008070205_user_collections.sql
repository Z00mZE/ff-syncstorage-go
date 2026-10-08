-- +goose Up
CREATE TABLE IF NOT EXISTS user_collections
(
    user_id       BIGINT      NOT NULL,
    collection_id INTEGER     NOT NULL,
    modified      TIMESTAMPTZ NOT NULL,
    count         BIGINT,
    total_bytes   BIGINT,
    PRIMARY KEY (user_id, collection_id)
);

-- +goose Down
DROP TABLE IF EXISTS user_collections;
