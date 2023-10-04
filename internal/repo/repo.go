package repo

import (
	"context"
	"database/sql"
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
		&p.Birthday,
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
		    $2;`

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
			&p.Birthday,
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

// TODO: add default values in sql migrations

func createEvent(ctx context.Context, tx *sql.Tx, personID uint64, eventType eventType) error {
	const q = `
		INSERT INTO person_event 
		    (
			person_id,
		    type,
		    status,
		    updated
		    )
		VALUES ($1, $2, $3,
		        now() at time zone 'utc'); -- updated`

	_, err := tx.ExecContext(ctx, q, personID, eventType, deferred)
	return err
}

func (r *Repo) transaction(ctx context.Context, eventType eventType, f func(tx *sql.Tx) (uint64, error)) error {
	const op = "Repo.transaction"

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("%s: can't start transaction: %w", op, err)
	}
	defer tx.Rollback() // TODO: handing error?

	personID, err := f(tx)
	if err != nil {
		return err
	}

	err = createEvent(ctx, tx, personID, eventType)
	if err != nil {
		return fmt.Errorf("%s: can't create event: %w", op, err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("%s: can't commit transaction: %w", op, err)
	}

	return nil
}

func (r *Repo) CreatePerson(ctx context.Context, p person) (uint64, error) {
	const op = "Repo.CreatePerson"
	const q = `
		INSERT INTO	person (
			first_name,
			middle_name,
			last_name, 
			birthday,
			sex,
			education,
		    removed,
			created,
		    updated)
		VALUES
		    ($1, $2, $3, $4, $5, $6, 
		     false, -- removed
		     now() at time zone 'utc', -- created
		     now() at time zone 'utc') -- updated
		RETURNING
			person_id;`

	var personID uint64

	err := r.transaction(ctx, created, func(tx *sql.Tx) (uint64, error) {
		row := tx.QueryRowContext(ctx, q,
			p.FirstName,
			p.MiddleName,
			p.LastName,
			p.Birthday,
			p.Sex,
			p.Education,
		)

		err := row.Scan(&personID)
		if err != nil {
			return 0, fmt.Errorf("can't create person: %w", err)
		}

		return personID, nil
	})

	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return personID, nil
}

// UpdatePerson returns true if the record with id found and false if the record does
// not exist or an error occurred.
func (r *Repo) UpdatePerson(ctx context.Context, personID uint64, p person) (bool, error) {
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
		    person_id = $7;`

	var count int64

	err := r.transaction(ctx, updated, func(tx *sql.Tx) (uint64, error) {
		res, err := tx.ExecContext(ctx, q,
			p.FirstName,
			p.MiddleName,
			p.LastName,
			p.Birthday,
			p.Sex,
			p.Education,
			personID,
		)
		if err != nil {
			return 0, fmt.Errorf("can't update: %w", err)
		}

		count, err = res.RowsAffected()
		if err != nil {
			return 0, fmt.Errorf("can't get query result: %w", err)
		}

		return personID, nil
	})

	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return count != 0, nil
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
		    person_id = $1 AND NOT removed;`

	var count int64

	err := r.transaction(ctx, removed, func(tx *sql.Tx) (uint64, error) {
		res, err := r.db.ExecContext(ctx, q, personID)
		if err != nil {
			return 0, fmt.Errorf("can't delete: %w", err)
		}

		count, err = res.RowsAffected()
		if err != nil {
			return 0, fmt.Errorf("can't get result: %w", err)
		}

		return personID, nil
	})

	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return count != 0, nil
}
