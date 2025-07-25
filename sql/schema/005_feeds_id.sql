-- +goose Up
ALTER TABLE feeds ADD COLUMN id UUID UNIQUE NOT NULL;

-- +goose Down
ALTER TABLE feeds DROP COLUMN id;