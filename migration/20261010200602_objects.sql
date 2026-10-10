-- +goose Up
CREATE TABLE if not exists objects
(
    id            INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id       INTEGER NOT NULL,
    collection_id INTEGER NOT NULL,
    bso_id        TEXT    NOT NULL,
    sort_index    INTEGER default 0,
    payload       TEXT    NOT NULL,
    modified      INTEGER NOT NULL,
    expiry        INTEGER NOT NULL
);

-- +goose Down
drop table if exists objects;