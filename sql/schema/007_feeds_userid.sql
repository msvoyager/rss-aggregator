-- +goose Up

ALTER TABLE feeds
    ALTER COLUMN user_id SET NOT NULL;


-- +goose Down
ALTER TABLE feeds 
    DROP COLUMN user_id;