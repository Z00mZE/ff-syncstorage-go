-- +goose Up
CREATE TABLE if not exists user_collections
(
    id            INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id       INTEGER  NOT NULL,
    collection_id INTEGER  NOT NULL,
    modified      DATETIME NOT NULL,
    count         INTEGER,
    total_bytes   INTEGER
);


-- +goose Down
drop table if exists user_collections;