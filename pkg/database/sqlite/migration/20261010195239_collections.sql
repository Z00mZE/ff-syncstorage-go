-- +goose Up
INSERT INTO collections (id, name)
VALUES (1, 'clients'),
       (2, 'crypto'),
       (3, 'forms'),
       (4, 'history'),
       (5, 'keys'),
       (6, 'meta'),
       (7, 'bookmarks'),
       (8, 'prefs'),
       (9, 'tabs'),
       (10, 'passwords'),
       (11, 'addons'),
       (12, 'addresses'),
       (13, 'creditcards');
UPDATE sqlite_sequence
SET seq = 100
WHERE name = 'collections';

-- +goose Down
DELETE FROM collections WHERE 1=1;
