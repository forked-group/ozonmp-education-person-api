-- +goose Up
CREATE TABLE IF NOT EXISTS person (
    person_id   bigserial   PRIMARY KEY,
    first_name  varchar(50) NOT NULL,
    middle_name varchar(50) NOT NULL,
    last_name   varchar(50) NOT NULL,
    birthday    timestamp   NOT NULL,
    sex         smallint    NOT NULL,
    education   smallint    NOT NULL,
    removed     boolean     NOT NULL,
    created     timestamp   NOT NULL,
    updated     timestamp   NOT NULL
);

CREATE TABLE IF NOT EXISTS person_event (
    event_id    bigserial    PRIMARY KEY,
    person_id   bigint       NOT NULL,
    type        smallint     NOT NULL,
    status      smallint     NOT NULL,
    payload     jsonb,
    updated		TIMESTAMP    NOT NULL,
    CONSTRAINT fk_person_id
        FOREIGN key (person_id) REFERENCES person (person_id)
);


-- +goose Down
-- DROP TABLE person_event;
-- DROP TABLE person;
