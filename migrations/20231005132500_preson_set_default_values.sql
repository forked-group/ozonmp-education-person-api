-- +goose Up
ALTER TABLE person
    ALTER COLUMN removed SET DEFAULT false,
    ALTER COLUMN created SET DEFAULT (now() at time zone 'utc'),
    ALTER COLUMN updated SET DEFAULT (now() at time zone 'utc');

ALTER TABLE person_event
    RENAME COLUMN event_id TO person_event_id;

ALTER TABLE person_event
    ALTER COLUMN status  SET DEFAULT 1, -- Deferred
    ALTER COLUMN updated SET DEFAULT (now() at time zone 'utc');

-- +goose Down
ALTER TABLE person_event
    ALTER COLUMN status  DROP DEFAULT,
    ALTER COLUMN updated DROP DEFAULT;

ALTER TABLE person
    ALTER COLUMN removed DROP DEFAULT,
    ALTER COLUMN created DROP DEFAULT,
    ALTER COLUMN updated DROP DEFAULT;

ALTER TABLE person_event
    RENAME COLUMN person_event_id TO event_id;
