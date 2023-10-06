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

type PersonEventRepo struct {
	db *sqlx.DB
}

func NewEventRepo(db *sqlx.DB) *PersonEventRepo {
	return &PersonEventRepo{
		db: db,
	}
}

func (r PersonEventRepo) transaction(ctx context.Context, f func(tx *sql.Tx) error) error {
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

func (r PersonEventRepo) Lock(ctx context.Context, n uint64) ([]model.PersonEvent, error) {
	const op = "PersonEventRepo.Lock"

	const q = `
		WITH
		    batch(person_event_id) AS (
				SELECT
					MIN(person_event_id)
				FROM
					person_event
				WHERE 
					status = $1/*deferred*/ AND
					person_id NOT IN (
						SELECT DISTINCT 
							person_id
						FROM
							person_event
						WHERE 
							status = $2/*processed*/	    		
					)
				GROUP BY 
				    person_id
				LIMIT
					$3/*n*/  
		    )
		UPDATE
			person_event AS p
		SET
			status = $2/*processed*/,
			updated = (now() at time zone 'utc')
		FROM
		    batch AS b
		WHERE
		    p.person_event_id = b.person_event_id
		RETURNING
			p.person_event_id,
		    person_id,
		    type,
		    status,
		    payload;
`
	var (
		events []model.PersonEvent
	)

	err := r.transaction(ctx, func(tx *sql.Tx) error {

		rows, err := tx.QueryContext(ctx, q, model.Deferred, model.Processed, n)
		if err != nil {
			return fmt.Errorf("can't update person_event: %w", err)
		}

		for rows.Next() {
			var (
				id          uint64
				eventType   model.EventType
				eventStatus model.EventStatus
				personID    uint64
				payload     []byte
				entity      *model.Person
			)

			err = rows.Scan(&id, &personID, &eventType, &eventStatus, &payload)
			if err != nil {
				return fmt.Errorf("can't scan row: %w", err)
			}

			if len(payload) == 0 {
				entity = &model.Person{ID: personID}

			} else {
				err = json.Unmarshal(payload, entity)
				if err != nil {
					return fmt.Errorf("can't unmarhal event payload: %w", err)
				}
			}

			events = append(events, model.PersonEvent{
				ID:     id,
				Type:   eventType,
				Status: eventStatus,
				Entity: entity,
			})
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return events, nil
}

func (r PersonEventRepo) Unlock(ctx context.Context, eventIDs []uint64) (uint64, error) {
	const op = "PersonEventRepo.Unlock"

	const q = `
		UPDATE
		    person_event
		SET
		    status = $1
		WHERE
		    person_event_id IN (%s);
`
	if len(eventIDs) == 0 {
		return 0, nil
	}

	qr := fmt.Sprintf(q, placeholderList(2, len(eventIDs)))
	lo.Debug("%s: qr: %v", op, qr)

	args := make([]any, 0, len(eventIDs)+1)
	args = append(args, model.Deferred)
	args = append(args, anySlice(eventIDs)...)

	var n int64

	err := r.transaction(ctx, func(tx *sql.Tx) error {

		res, err := r.db.ExecContext(ctx, qr, args...)
		if err != nil {
			return fmt.Errorf("%s: can't update person_event: %w", op, err)
		}

		n, err = res.RowsAffected()
		if err != nil {
			return fmt.Errorf("can't get number of removed rows: %w", err)
		}

		return err
	})
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return uint64(n), nil
}

func (r PersonEventRepo) Remove(ctx context.Context, eventIDs []uint64) (uint64, error) {
	const op = "PersonEventRepo.RemovePerson"
	const q = `
		DELETE FROM
		    person_event
		WHERE
		    person_event_id IN (%s);
`
	if len(eventIDs) == 0 {
		return 0, nil
	}

	qr := fmt.Sprintf(q, placeholderList(1, len(eventIDs)))
	args := anySlice(eventIDs)

	var n int64

	err := r.transaction(ctx, func(tx *sql.Tx) error {

		res, err := tx.ExecContext(ctx, qr, args...)
		if err != nil {
			return fmt.Errorf("can't update person_event: %w", err)
		}

		n, err = res.RowsAffected()
		if err != nil {
			return fmt.Errorf("can't get number of removed rows: %w", err)
		}

		return err
	})
	if err != nil {
		err = fmt.Errorf("%s: %w", op, err)
	}

	return uint64(n), nil
}
