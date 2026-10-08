-- +goose Up
CREATE TABLE IF NOT EXISTS basic_storage_objects (
                      user_id BIGINT NOT NULL,
                      collection_id INTEGER NOT NULL,
                      bso_id TEXT NOT NULL,
                      sortindex INTEGER,
                      payload TEXT NOT NULL,
                      modified TIMESTAMPTZ NOT NULL,
                      expiry TIMESTAMPTZ NOT NULL,
                      PRIMARY KEY (
                                   user_id,
                                   collection_id,
                                   bso_id
                          ),
                      FOREIGN KEY (user_id, collection_id) REFERENCES user_collections (user_id, collection_id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS basic_storage_objects;
