package repo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Repo struct {
	db        *sqlx.DB
	batchSize uint
}

// NewRepo returns Repo interface
func NewRepo(db *sqlx.DB, batchSize uint) *Repo {
	return &Repo{
		db:        db,
		batchSize: batchSize,
	}
}

func (r *Repo) DescribePerson(ctx context.Context, personID uint64) (*person, error) {
	const op = "Repo.DescribePerson"

	const q = `
		SELECT
			first_name,
			middle_name,
			last_name,
			birthday,
			sex,
			education,
			removed,
			created,
			updated
		FROM
		    person
		WHERE
		    person_id = $1 AND NOT removed;`

	p := &person{ID: personID}

	err := r.db.QueryRowContext(ctx, q, personID).Scan(
		&p.FirstName,
		&p.MiddleName,
		&p.LastName,
		&p.Birthday.Time,
		&p.Sex,
		&p.Education,
		&p.Removed,
		&p.Created,
		&p.Updated,
	)
	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("%s: can't select: %w", op, err)
	}

	return p, nil
}

func (r *Repo) ListPerson(ctx context.Context, cursor uint64, limit uint64) ([]person, error) {
	const op = "Repo.ListPerson"

	const q = `
		SELECT
			person_id,
			first_name,
			middle_name,
			last_name,
			birthday,
			sex,
			education,
			removed,
			created,
			updated
		FROM
			person
		WHERE
			person_id >= $1 AND NOT removed
		ORDER BY
			person_id
		LIMIT
			$2;
`

	if limit <= 0 || limit > uint64(r.batchSize) {
		limit = uint64(r.batchSize)
	}

	rows, err := r.db.QueryContext(ctx, q, cursor, limit)
	if err != nil {
		return nil, fmt.Errorf("%s: can't select: %w", op, err)
	}

	list := make([]person, 0, limit)

	for rows.Next() {
		var p person
		err = rows.Scan(
			&p.ID,
			&p.FirstName,
			&p.MiddleName,
			&p.LastName,
			&p.Birthday.Time,
			&p.Sex,
			&p.Education,
			&p.Removed,
			&p.Created,
			&p.Updated,
		)
		if err != nil {
			return list, fmt.Errorf("%s: can't scan row: %w", op, err)
		}

		list = append(list, p)
	}

	return list, nil
}

func createEvent(ctx context.Context, tx *sql.Tx, eventType eventType, person *person) error {
	const op = "repo.createEvent"

	const q = `
		INSERT INTO
		    person_event
		    (
		     person_id,
		     type,
		     payload
		    )
		VALUES ($1, $2, $3);
`

	payload, err := json.Marshal(person)
	if err != nil {
		return fmt.Errorf("%s: can't marshal person: %w", op, err)
	}

	_, err = tx.ExecContext(ctx, q, person.ID, eventType, payload)
	return err
}

func (r *Repo) transaction(ctx context.Context, f func(tx *sql.Tx) error) error {
	const op = "Repo.transaction"

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("%s: can't start transaction: %w", op, err)
	}
	defer tx.Rollback() // TODO: handing error?

	err = f(tx)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("%s: can't commit transaction: %w", op, err)
	}

	return nil
}

func (r *Repo) CreatePerson(ctx context.Context, pc personCreate) (uint64, error) {
	const op = "Repo.CreatePerson"

	const q = `
		INSERT INTO	person
		    (
			 first_name,
			 middle_name,
			 last_name, 
			 birthday,
			 sex,
			 education
			)
		VALUES
		    ($1, $2, $3, $4, $5, $6) 
		RETURNING
			person_id,
			created;
`
	var p person

	err := r.transaction(ctx, func(tx *sql.Tx) error {

		row := tx.QueryRowContext(ctx, q,
			pc.FirstName,
			pc.MiddleName,
			pc.LastName,
			pc.Birthday.Time,
			pc.Sex,
			pc.Education,
		)

		err := row.Scan(&p.ID, &p.Created)
		if err != nil {
			return fmt.Errorf("can't create person: %w", err)
		}

		p.PersonCreate = pc

		err = createEvent(ctx, tx, created, &p)
		if err != nil {
			return fmt.Errorf("can't create event: %w", err)
		}

		return nil
	})

	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return p.ID, nil
}

// UpdatePerson returns true if the record with id found and false if the record does
// not exist or an error occurred.
func (r *Repo) UpdatePerson(ctx context.Context, personID uint64, pc personCreate) (bool, error) {
	const op = "Repo.UpdatePerson"

	const q = `
		UPDATE
		    person
		SET 
		    first_name 	= $1,
			middle_name = $2,
			last_name 	= $3,
			birthday    = $4,
			sex 		= $5,
			education   = $6,
			updated     = (now() at time zone 'utc')
		WHERE
		    person_id = $7 AND NOT removed 
		RETURNING
			updated;
`
	var ok bool

	err := r.transaction(ctx, func(tx *sql.Tx) error {

		row := tx.QueryRowContext(ctx, q,
			pc.FirstName,
			pc.MiddleName,
			pc.LastName,
			pc.Birthday.Time,
			pc.Sex,
			pc.Education,
			personID,
		)

		var p person

		err := row.Scan(&p.Updated)
		if err == sql.ErrNoRows {
			return nil
		}
		if err != nil {
			return fmt.Errorf("can't update: %w", err)
		}

		ok = true
		p.ID = personID
		p.PersonCreate = pc

		err = createEvent(ctx, tx, updated, &p)
		if err != nil {
			return fmt.Errorf("can't create event: %w", err)
		}

		return nil
	})

	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return ok, nil
}

// RemovePerson returns true if the record was indeed removed, and false if the
// record does not exist or an error occurred.
func (r *Repo) RemovePerson(ctx context.Context, personID uint64) (bool, error) {
	const op = "Repo.RemovePerson"

	const q = `
		UPDATE 
		    person 
		SET
		    removed = true,
		    updated = (now() at time zone 'utc')
		WHERE
		    person_id = $1 AND NOT removed
		RETURNING
			updated;
`
	var person person

	err := r.transaction(ctx, func(tx *sql.Tx) error {

		row := tx.QueryRowContext(ctx, q, personID)

		err := row.Scan(&person.Updated)
		if err == sql.ErrNoRows {
			return nil
		}
		if err != nil {
			return fmt.Errorf("can't update: %w", err)
		}

		err = createEvent(ctx, tx, removed, &person)
		if err != nil {
			return fmt.Errorf("can't create event: %w", err)
		}

		return nil
	})

	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return person.Removed, nil
}
