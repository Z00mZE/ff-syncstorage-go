-- +goose Up
CREATE TABLE if not exists batches
(
    id            INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id       INTEGER NOT NULL,
    collection_id INTEGER NOT NULL,
    batch_id      TEXT    NOT NULL,
    expiry        INTEGER NOT NULL
);

-- +goose Down
drop table if exists batches;