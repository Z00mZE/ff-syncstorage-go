-- +goose Up
create unique index if not exists batches_uid_cid_bid_uniq
    on batches (user_id, collection_id, batch_id);

-- +goose Down
drop index if exists batches_uid_cid_bid_uniq;
