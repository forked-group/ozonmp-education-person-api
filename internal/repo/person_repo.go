package repo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	model "github.com/aaa2ppp/ozonmp-education-person-api/internal/model/education"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
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

func (r *PersonRepo) DescribePerson(ctx context.Context, personID uint64) (*model.Person, error) {
	const op = "PersonRepo.DescribePerson"

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

	p := &model.Person{ID: personID}

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
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("%s: can't select: %w", op, err)
	}

	return p, nil
}

func (r *PersonRepo) ListPerson(ctx context.Context, cursor uint64, limit uint64) ([]model.Person, error) {
	const op = "PersonRepo.ListPerson"

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

	list := make([]model.Person, 0, limit)

	for rows.Next() {
		var p model.Person
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

func createPayload(p model.Person, f model.PersonField) ([]byte, error) {
	pj := &struct {
		ID         uint64          `json:"id,omitempty"`
		FirstName  string          `json:"first_name,omitempty"`
		MiddleName string          `json:"middle_name,omitempty"`
		LastName   string          `json:"last_name,omitempty"`
		Birthday   *time.Time      `json:"birthday,omitempty"`
		Sex        model.Sex       `json:"sex,omitempty"`
		Education  model.Education `json:"education,omitempty"`
		Removed    bool            `json:"removed,omitempty"`
		Created    *time.Time      `json:"created,omitempty"`
		Updated    *time.Time      `json:"updated,omitempty"`
	}{}

	if f.IsSet(model.PersonID) {
		pj.ID = p.ID
	}
	if f.IsSet(model.PersonFirstName) {
		pj.FirstName = p.FirstName
	}
	if f.IsSet(model.PersonMiddleName) {
		pj.MiddleName = p.MiddleName
	}
	if f.IsSet(model.PersonLastName) {
		pj.LastName = p.LastName
	}
	if f.IsSet(model.PersonBirthday) {
		pj.Birthday = &p.Birthday
	}
	if f.IsSet(model.PersonSex) {
		pj.Sex = p.Sex
	}
	if f.IsSet(model.PersonEducation) {
		pj.Education = p.Education
	}
	if f.IsSet(model.PersonRemoved) {
		pj.Removed = p.Removed
	}
	if f.IsSet(model.PersonCreated) {
		pj.Created = &p.Created
	}
	if f.IsSet(model.PersonUpdated) {
		pj.Updated = &p.Updated
	}

	return json.Marshal(pj)
}

func createEvent(ctx context.Context, tx *sql.Tx, eventType model.EventType, person model.Person, fields model.PersonField) error {
	const op = "repo.createEvent"

	const q = `
		INSERT INTO person_event 
		    (person_id, type, payload)	
		VALUES 
		    ($1, $2, $3);
`

	payload, err := createPayload(person, fields)
	if err != nil {
		return fmt.Errorf("%s: can't create payload: %w", op, err)
	}

	_, err = tx.ExecContext(ctx, q, person.ID, eventType, payload)
	return err
}

func (r *PersonRepo) transaction(ctx context.Context, f func(tx *sql.Tx) error) error {
	const op = "PersonRepo.transaction"

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

const editableFields = model.PersonFirstName |
	model.PersonMiddleName |
	model.PersonLastName |
	model.PersonBirthday |
	model.PersonSex |
	model.PersonEducation

func (r *PersonRepo) CreatePerson(ctx context.Context, person model.Person) (uint64, error) {
	const op = "PersonRepo.CreatePerson"

	const q = `
		INSERT INTO person 
			(%s)
		VALUES 
			(%s)
		RETURNING 
			person_id,
			created,
			updated;
`
	p := newParamsWithMask(editableFields)

	p.add(model.PersonFirstName, "first_name", person.FirstName)
	p.add(model.PersonMiddleName, "middle_name", person.MiddleName)
	p.add(model.PersonLastName, "last_name", person.LastName)
	p.add(model.PersonBirthday, "birthday", person.Birthday)
	p.add(model.PersonSex, "sex", person.Sex)
	p.add(model.PersonEducation, "education", person.Education)

	p.fields |= model.PersonID
	p.fields |= model.PersonCreated
	p.fields |= model.PersonUpdated

	qr := fmt.Sprintf(q, strings.Join(p.names, ","), placeholderList(1, len(p.names)))

	err := r.transaction(ctx, func(tx *sql.Tx) error {
		row := tx.QueryRowContext(ctx, qr, p.values...)

		err := row.Scan(&person.ID, &person.Created, &person.Updated)
		if err != nil {
			return fmt.Errorf("can't create person: %w", err)
		}

		err = createEvent(ctx, tx, model.Created, person, p.fields)
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

// UpdatePerson returns true if the record with id found and false if the record does
// not exist or an error occurred.
func (r *PersonRepo) UpdatePerson(ctx context.Context, personID uint64, person model.Person, fields model.PersonField) (bool, error) {
	const op = "PersonRepo.UpdatePerson"

	const q = `
		UPDATE 
			person
		SET 
			%s,	updated = (now() at time zone 'utc')
		WHERE
			person_id=$1 AND NOT removed 
		RETURNING
			updated;
`
	person.ID = personID
	p := newParamsWithMask(fields & editableFields)

	p.values = append(p.values, person.ID) // $1

	p.add(model.PersonFirstName, "first_name", person.FirstName)
	p.add(model.PersonMiddleName, "middle_name", person.MiddleName)
	p.add(model.PersonLastName, "last_name", person.LastName)
	p.add(model.PersonBirthday, "birthday", person.Birthday)
	p.add(model.PersonSex, "sex", person.Sex)
	p.add(model.PersonEducation, "education", person.Education)

	p.fields |= model.PersonID
	p.fields |= model.PersonUpdated

	qr := fmt.Sprintf(q, placeholderSetList(2, p.names)) // $2...

	var ok bool

	err := r.transaction(ctx, func(tx *sql.Tx) error {

		row := tx.QueryRowContext(ctx, qr, p.values...)

		err := row.Scan(&person.Updated)
		if err == sql.ErrNoRows {
			return nil
		}
		if err != nil {
			return fmt.Errorf("can't update: %w", err)
		}

		ok = true

		err = createEvent(ctx, tx, model.Updated, person, p.fields)
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
func (r *PersonRepo) RemovePerson(ctx context.Context, personID uint64) (bool, error) {
	const op = "PersonRepo.RemovePerson"

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
	person := model.Person{ID: personID, Removed: true}
	fields := model.PersonID | model.PersonRemoved | model.PersonUpdated

	var ok bool

	err := r.transaction(ctx, func(tx *sql.Tx) error {

		row := tx.QueryRowContext(ctx, q, person.ID)

		err := row.Scan(&person.Updated)
		if err == sql.ErrNoRows {
			return nil
		}
		if err != nil {
			return fmt.Errorf("can't update: %w", err)
		}

		ok = true

		err = createEvent(ctx, tx, model.Removed, person, fields)
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
