-- +goose Up
CREATE TABLE if not exists collections
(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name          TEXT NOT NULL UNIQUE
);

-- +goose Down
drop table if exists collections;
