-- +goose Up
ALTER TABLE person
    ADD CONSTRAINT name_not_empty CHECK (first_name != '' OR middle_name != '' OR last_name != '');

-- +goose Down
ALTER TABLE person
    DROP CONSTRAINT name_not_empty;
