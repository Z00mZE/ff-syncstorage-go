-- +goose Up
create unique index if not exists objects_cid_uid_oid_uniq
    on objects (user_id, collection_id, bso_id);

-- +goose Down
drop index if exists objects_cid_uid_oid_uniq;
