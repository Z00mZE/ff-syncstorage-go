-- +goose Up
CREATE TABLE batch_objects
(
    id              INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id         INTEGER NOT NULL,
    collection_id   INTEGER NOT NULL,
    batch_id        TEXT    NOT NULL,
    batch_object_id TEXT    NOT NULL,
    sort_index      INTEGER,
    payload         TEXT,
    ttl             INTEGER
);

-- +goose Down
drop table if exists batch_objects;
