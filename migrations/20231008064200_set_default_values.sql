-- +goose Up
ALTER TABLE person
    ALTER COLUMN first_name  SET DEFAULT '',
    ALTER COLUMN middle_name SET DEFAULT '',
    ALTER COLUMN last_name   SET DEFAULT '',
    ALTER COLUMN sex         SET DEFAULT 0,
    ALTER COLUMN education   SET DEFAULT 0;

ALTER TABLE person_event
    ALTER COLUMN type   SET DEFAULT 0,
    ALTER COLUMN status SET DEFAULT 0;

-- +goose Down
ALTER TABLE person
    ALTER COLUMN first_name  DROP DEFAULT,
    ALTER COLUMN middle_name DROP DEFAULT,
    ALTER COLUMN last_name   DROP DEFAULT,
    ALTER COLUMN sex         DROP DEFAULT,
    ALTER COLUMN education   DROP DEFAULT;

ALTER TABLE person_event
    ALTER COLUMN type   DROP DEFAULT,
    ALTER COLUMN status SET  DEFAULT 1;
