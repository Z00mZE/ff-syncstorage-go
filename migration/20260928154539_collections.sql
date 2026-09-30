-- +goose Up
CREATE TABLE IF NOT EXISTS collections
(
    collection_id serial PRIMARY KEY,
    name          VARCHAR(32) NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE IF EXISTS collections;
