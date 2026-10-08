-- +goose Up

CREATE TABLE batch_basic_storage_objects
(
    user_id       BIGINT  NOT NULL,
    collection_id INTEGER NOT NULL,
    batch_id      UUID    NOT NULL,
    batch_bso_id  TEXT    NOT NULL,
    sort_index     INTEGER,
    payload       TEXT,
    ttl           BIGINT,
    PRIMARY KEY (user_id, collection_id, batch_id, batch_bso_id),
    FOREIGN KEY (user_id, collection_id, batch_id)
        REFERENCES batches (user_id, collection_id, batch_id)
        ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS batch_basic_storage_objects;
