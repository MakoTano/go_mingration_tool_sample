-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO book (category_id,title) VALUES (10,'サルの山脈');

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back


