-- +goose Up
create unique index if not exists batch_objects_uid_cid_bid_uniq
    on batch_objects (user_id, collection_id, batch_id, batch_object_id);

-- +goose Down
drop index if exists batch_objects_uid_cid_bid_uniq;
