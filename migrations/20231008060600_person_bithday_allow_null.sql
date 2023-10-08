-- +goose Up
ALTER TABLE person
    ALTER COLUMN birthday DROP NOT NULL;

-- +goose Down
ALTER TABLE person
    ALTER COLUMN birthday SET NOT NULL;