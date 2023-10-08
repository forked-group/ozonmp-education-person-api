package repo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
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

func (r *PersonRepo) DescribePerson(ctx context.Context, personID uint64) (*model.Person, error) {
	const op = "PersonRepo.DescribePerson"

	const query = `
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
		    person_id = $1 AND NOT removed;
`
	person := &model.Person{ID: personID}

	row := r.db.QueryRowContext(ctx, query, person.ID)

	var birthday sql.NullTime
	err := row.Scan(
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
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("%s: can't select: %w", op, err)
	}

	if birthday.Valid {
		person.Birthday = &model.Date{Time: birthday.Time}
	}

	return person, nil
}

func (r *PersonRepo) ListPerson(ctx context.Context, cursor uint64, limit uint64) ([]model.Person, error) {
	const op = "PersonRepo.ListPerson"

	const query = `
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

	rows, err := r.db.QueryContext(ctx, query, cursor, limit)
	if err != nil {
		return nil, fmt.Errorf("%s: can't select: %w", op, err)
	}

	list := make([]model.Person, 0, limit)

	for rows.Next() {
		var person model.Person
		var birthday sql.NullTime

		err = rows.Scan(
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
			return list, fmt.Errorf("%s: can't scan row: %w", op, err)
		}

		if birthday.Valid {
			person.Birthday = &model.Date{Time: birthday.Time}
		}

		list = append(list, person)
	}

	return list, nil
}

func createEventPayload(person model.Person, fields model.PersonField) ([]byte, error) {
	f := FieldSet[model.PersonField](fields)
	m := map[string]any{}

	if f.Includes(model.PersonID) {
		m["person_id"] = person.ID
	}
	if f.Includes(model.PersonFirstName) {
		m["first_name"] = person.FirstName
	}
	if f.Includes(model.PersonMiddleName) {
		m["middle_name"] = person.MiddleName
	}
	if f.Includes(model.PersonLastName) {
		m["last_name"] = person.LastName
	}
	if f.Includes(model.PersonBirthday) {
		m["birthday"] = &person.Birthday
	}
	if f.Includes(model.PersonSex) {
		m["sex"] = person.Sex
	}
	if f.Includes(model.PersonEducation) {
		m["education"] = person.Education
	}
	if f.Includes(model.PersonRemoved) {
		m["removed"] = person.Removed
	}
	if f.Includes(model.PersonCreated) {
		m["created"] = &person.Created
	}
	if f.Includes(model.PersonUpdated) {
		m["updated"] = &person.Updated
	}

	return json.Marshal(m)
}

func createEvent(ctx context.Context, tx *sql.Tx, eventType model.EventType, person model.Person, fields model.PersonField) error {
	const op = "repo.createEvent"

	const query = `
		INSERT INTO person_event 
		    (
		     person_id,
		     type,
		     status,
		     payload
		    )	
		VALUES 
		    ($1, $2, $3, $4);
`

	payload, err := createEventPayload(person, fields)
	if err != nil {
		return fmt.Errorf("%s: can't create payload: %w", op, err)
	}

	_, err = tx.ExecContext(ctx, query,
		person.ID,
		eventType,
		model.Deferred,
		payload,
	)
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

	var query = `
		INSERT INTO person 
			(%s)
		VALUES 
			(%s)
		RETURNING 
			person_id,
			created,
			updated;
`
	h := NewListsBuilder(editableFields)
	h.AddField(model.PersonFirstName, "first_name", person.FirstName)
	h.AddField(model.PersonMiddleName, "middle_name", person.MiddleName)
	h.AddField(model.PersonLastName, "last_name", person.LastName)
	h.AddField(model.PersonBirthday, "birthday", person.Birthday.NullTime())
	h.AddField(model.PersonSex, "sex", person.Sex)
	h.AddField(model.PersonEducation, "education", person.Education)

	fieldsToEvent := h.Fields |
		model.PersonID |
		model.PersonCreated |
		model.PersonUpdated

	query = fmt.Sprintf(query, h.NameList(), h.ValueListTemplate())

	err := r.transaction(ctx, func(tx *sql.Tx) error {
		row := tx.QueryRowContext(ctx, query, h.Args...)

		err := row.Scan(
			&person.ID,
			&person.Created,
			&person.Updated,
		)
		if err != nil {
			return fmt.Errorf("can't create person: %w", err)
		}

		err = createEvent(ctx, tx, model.Created, person, fieldsToEvent)
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

	var query = `
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

	b := NewListsBuilder(fields & editableFields)
	b.Args = []any{person.ID}
	b.AddField(model.PersonFirstName, "first_name", person.FirstName)
	b.AddField(model.PersonMiddleName, "middle_name", person.MiddleName)
	b.AddField(model.PersonLastName, "last_name", person.LastName)
	b.AddField(model.PersonBirthday, "birthday", person.Birthday.NullTime())
	b.AddField(model.PersonSex, "sex", person.Sex)
	b.AddField(model.PersonEducation, "education", person.Education)

	fieldsToEvent := b.Fields |
		model.PersonID |
		model.PersonUpdated

	query = fmt.Sprintf(query, b.FieldListTemplate())

	var ok bool

	err := r.transaction(ctx, func(tx *sql.Tx) error {
		row := tx.QueryRowContext(ctx, query, b.Args...)

		err := row.Scan(&person.Updated)
		if err == sql.ErrNoRows {
			return nil
		}
		if err != nil {
			return fmt.Errorf("can't update: %w", err)
		}

		ok = true

		err = createEvent(ctx, tx, model.Updated, person, fieldsToEvent)
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

	const query = `
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
	person := model.Person{
		ID:      personID,
		Removed: true,
	}

	fieldsToEvent := model.PersonID |
		model.PersonRemoved |
		model.PersonUpdated

	var ok bool

	err := r.transaction(ctx, func(tx *sql.Tx) error {
		row := tx.QueryRowContext(ctx, query, person.ID)

		err := row.Scan(&person.Updated)
		if err == sql.ErrNoRows {
			return nil
		}
		if err != nil {
			return fmt.Errorf("can't update: %w", err)
		}

		ok = true

		err = createEvent(ctx, tx, model.Removed, person, fieldsToEvent)
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
