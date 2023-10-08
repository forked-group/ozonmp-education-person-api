-- +goose Up
ALTER TABLE person
    ALTER COLUMN birthday TYPE date;

-- +goose Down
ALTER TABLE person
    ALTER COLUMN birthday TYPE timestamp;