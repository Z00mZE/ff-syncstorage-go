-- +goose Up
create unique index if not exists user_collections_collection_id_user_id_uindex
    on user_collections (collection_id, user_id);


-- +goose Down
drop index if exists user_collections_collection_id_user_id_uindex;
