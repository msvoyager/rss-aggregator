-- +goose Up
ALTER TABLE feeds
ADD PRIMARY KEY (id);

-- +goose Down
ALTER TABLE feeds
DROP PRIMARY KEY;