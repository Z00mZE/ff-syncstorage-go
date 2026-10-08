-- +goose Up
CREATE INDEX IF NOT EXISTS bsos_modified_idx ON basic_storage_objects (
                                                         user_id,
                                                         collection_id,
                                                         modified DESC
    );

-- +goose Down
DROP INDEX IF EXISTS bsos_modified_idx;
