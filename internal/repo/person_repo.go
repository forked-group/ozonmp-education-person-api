package repo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/lib/log/lo"
	model "github.com/aaa2ppp/ozonmp-education-person-api/internal/model/education"
	"github.com/jmoiron/sqlx"
)

type PersonRepo struct {
	db        *sqlx.DB
	batchSize uint
}

func NewPersonRepo(db *sqlx.DB, batchSize uint) *PersonRepo {
	return &PersonRepo{
		db:        db,
		batchSize: batchSize,
	}
}

var personRowFields = CommaList([]string{
	"person_id",
	"first_name",
	"middle_name",
	"last_name",
	"birthday",
	"sex",
	"education",
	"removed",
	"created",
	"updated",
})

func scanPerson(row interface{ Scan(...any) error }, person *model.Person) error {
	var birthday sql.NullTime

	err := row.Scan(
		&person.ID,
		&person.FirstName,
		&person.MiddleName,
		&person.LastName,
		&birthday,
		&person.Sex,
		&person.Education,
		&person.Removed,
		&person.Created,
		&person.Updated,
	)
	if err != nil {
		return err
	}

	if birthday.Valid {
		person.Birthday = &model.Date{Time: birthday.Time}
	}

	return nil
}

func (r *PersonRepo) Describe(ctx context.Context, personID uint64) (*model.Person, error) {
	const op = "PersonRepo.Describe"
	var person model.Person

	var query = `
		SELECT
		    /*personRowFields*/%s
		FROM
		    person
		WHERE
		    person_id = /*personID*/$1 AND NOT removed;
`
	query = fmt.Sprintf(query, personRowFields)

	row := r.db.QueryRowContext(ctx, query, personID)

	err := scanPerson(row, &person)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("%s: can't do query: %w", op, err)
	}

	return &person, nil
}

func (r *PersonRepo) List(ctx context.Context, cursor uint64, limit uint64) ([]model.Person, error) {
	const op = "PersonRepo.List"

	if limit <= 0 || limit > uint64(r.batchSize) {
		limit = uint64(r.batchSize)
	}

	list := make([]model.Person, 0, limit)

	var query = `
		SELECT
		    /*personRowFields*/%s
		FROM
			person
		WHERE
			person_id >= /*cursor*/$1 AND NOT removed
		ORDER BY
			person_id
		LIMIT
			/*limit*/$2;
`
	query = fmt.Sprintf(query, personRowFields)

	rows, err := r.db.QueryContext(ctx, query, cursor, limit)
	if err != nil {
		return nil, fmt.Errorf("%s: can't do query: %w", op, err)
	}

	for rows.Next() {
		var person model.Person

		err = scanPerson(rows, &person)
		if err != nil {
			return nil, fmt.Errorf("%s: can't scan row: %w", op, err)
		}

		list = append(list, person)
	}

	return list, nil
}

func createPersonRepoEvent(ctx context.Context, tx *sql.Tx, eventType model.EventType, person model.Person) error {
	const op = "createPersonRepoEvent"

	payload, err := json.Marshal(person)
	if err != nil {
		return fmt.Errorf("%s: can't marshal person: %w", op, err)
	}

	const query = `
		INSERT INTO person_event(
			person_id,
			type,
			status,
			payload
		)	
		VALUES 
		    ($1, $2, $3, $4);
`
	_, err = tx.ExecContext(ctx, query,
		person.ID,
		eventType,
		model.Deferred,
		payload,
	)
	if err != nil {
		return fmt.Errorf("%s: can't exec query: %w", op, err)
	}

	return nil
}

func (r *PersonRepo) transaction(ctx context.Context, f func(tx *sql.Tx) error) error {
	const op = "PersonRepo.transaction"

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("%s: can't start transaction: %w", op, err)
	}
	defer tx.Rollback() // TODO: error handing?

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

const editablePersonFields = model.PersonFirstName |
	model.PersonMiddleName |
	model.PersonLastName |
	model.PersonBirthday |
	model.PersonSex |
	model.PersonEducation

func addPersonFields(b *ListsBuilder, person model.Person, fields model.PersonField) {
	mask := FieldSet[model.PersonField](editablePersonFields & fields)

	b.AddFieldIf(mask.Includes(model.PersonFirstName), "first_name", person.FirstName)
	b.AddFieldIf(mask.Includes(model.PersonMiddleName), "middle_name", person.MiddleName)
	b.AddFieldIf(mask.Includes(model.PersonLastName), "last_name", person.LastName)
	b.AddFieldIf(mask.Includes(model.PersonBirthday), "birthday", person.Birthday.NullTime())
	b.AddFieldIf(mask.Includes(model.PersonSex), "sex", person.Sex)
	b.AddFieldIf(mask.Includes(model.PersonEducation), "education", person.Education)
}

func (r *PersonRepo) Create(ctx context.Context, person model.Person) (uint64, error) {
	const op = "PersonRepo.Create"

	var query = `
		INSERT INTO person 
			(/*NameList*/%s)
		VALUES 
			(/*ValueListTemplate*/%s)
		RETURNING
			/*personRowFields*/%s; 
`
	var b ListsBuilder
	addPersonFields(&b, person, editablePersonFields)

	query = fmt.Sprintf(query, b.NameList(), b.ValueListTemplate(), personRowFields)
	lo.Debug("%s: query: %s", op, query)

	err := r.transaction(ctx, func(tx *sql.Tx) error {
		row := tx.QueryRowContext(ctx, query, b.Args...)

		err := scanPerson(row, &person)
		if err != nil {
			return fmt.Errorf("can't do query: %w", err)
		}

		err = createPersonRepoEvent(ctx, tx, model.Created, person)
		if err != nil {
			return fmt.Errorf("can't create event: %w", err)
		}

		return nil
	})

	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return person.ID, nil
}

// Update returns true if the record with id found and false if the record does
// not exist or an error occurred.
func (r *PersonRepo) Update(ctx context.Context, personID uint64, person model.Person, fields model.PersonField) (bool, error) {
	const op = "PersonRepo.Update"

	fields &= editablePersonFields
	if fields == 0 {
		return false, fmt.Errorf("%s: no any fields to update", op)
	}

	var query = `
		UPDATE 
			person
		SET 
			/*FieldListTemplate*/%s,
		    updated = (now() at time zone 'utc')
		WHERE
			person_id = /*personID*/$1 AND NOT removed 
		RETURNING
			/*personRowFields*/%s;
`
	var b ListsBuilder
	b.Args = []any{personID}
	addPersonFields(&b, person, fields)

	query = fmt.Sprintf(query, b.FieldListTemplate(), personRowFields)
	lo.Debug("%s: query: %s", op, query)

	var ok bool

	err := r.transaction(ctx, func(tx *sql.Tx) error {
		row := tx.QueryRowContext(ctx, query, b.Args...)

		err := scanPerson(row, &person)
		if err == sql.ErrNoRows {
			return nil
		}
		if err != nil {
			return fmt.Errorf("can't do query: %w", err)
		}

		ok = true

		err = createPersonRepoEvent(ctx, tx, model.Updated, person)
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

// Remove returns true if the record was indeed removed, and false if the
// record does not exist or an error occurred.
func (r *PersonRepo) Remove(ctx context.Context, personID uint64) (bool, error) {
	const op = "PersonRepo.Remove"

	var query = `
		UPDATE 
		    person
		SET 
		    removed = true,
		    updated = (now() at time zone 'utc')
		WHERE
		    person_id = /*personID*/$1 AND NOT removed
		RETURNING
			/*personRowFields*/%s;
`
	query = fmt.Sprintf(query, personRowFields)
	lo.Debug("%s: query: %s", op, query)

	var ok bool

	err := r.transaction(ctx, func(tx *sql.Tx) error {
		row := tx.QueryRowContext(ctx, query, personID)

		var person model.Person
		err := scanPerson(row, &person)
		if err == sql.ErrNoRows {
			return nil
		}
		if err != nil {
			return fmt.Errorf("can't do query: %w", err)
		}

		ok = true

		err = createPersonRepoEvent(ctx, tx, model.Updated, person)
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
