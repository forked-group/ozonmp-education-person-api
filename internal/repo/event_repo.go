package repo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/lib/log/lo"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/model/education"
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
)

type EventRepo struct {
	db *sqlx.DB
}

func NewEventRepo(db *sqlx.DB) *EventRepo {
	return &EventRepo{
		db: db,
	}
}

func (r EventRepo) transaction(ctx context.Context, f func(tx *sql.Tx) error) error {
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

func (r EventRepo) Lock(ctx context.Context, n uint64) ([]education.PersonEvent, error) {
	const op = "EventRepo.Lock"
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
	var events []personEvent

	err := r.transaction(ctx, func(tx *sql.Tx) error {

		rows, err := tx.QueryContext(ctx, q, deferred, processed, n)
		if err != nil {
			return fmt.Errorf("can't update person_event: %w", err)
		}

		for rows.Next() {
			var (
				id       uint64
				type_    eventType
				status   eventStatus
				personID uint64
				payload  []byte
				entity   *person
			)

			err = rows.Scan(&id, &personID, &type_, &status, &payload)
			if err != nil {
				return fmt.Errorf("can't scan row: %w", err)
			}

			if len(payload) == 0 {
				entity = &person{ID: personID}

			} else {
				err = json.Unmarshal(payload, entity)
				if err != nil {
					return fmt.Errorf("can't unmarhal event payload: %w", err)
				}
			}

			events = append(events, personEvent{
				ID:     id,
				Type:   type_,
				Status: status,
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

func anySlice[T any](a []T) []any {
	res := make([]any, len(a))

	for i := range res {
		res[i] = a[i]
	}

	return res
}

func placeholderList(first int, n int) string {
	elems := make([]string, n)

	for i, j := 0, first; i < n; i, j = i+1, j+1 {
		elems[i] = "$" + strconv.Itoa(j)
	}

	return strings.Join(elems, ",")
}

func (r EventRepo) Unlock(ctx context.Context, eventIDs []uint64) (uint64, error) {
	const op = "EventRepo.Unlock"
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
	args = append(args, deferred)
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

func (r EventRepo) Remove(ctx context.Context, eventIDs []uint64) (uint64, error) {
	const op = "EventRepo.Remove"
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
