package consumer

import (
	"context"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/lib/log/lo"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/lib/sync/sy"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/model/person"
	"time"
)

//go:generate mockgen -destination=./mocks/event_locker.go . EventLocker
type EventLocker interface {
	Lock(n uint64) ([]person.PersonEvent, error)
}

type SuccessFunc func()
type IncompleteFunc func()
type FailFunc func(err error)

type Consumer struct {
	events    chan *person.PersonEvent
	repo      EventLocker
	batchSize int
	timeout   time.Duration
}

func New(
	events chan *person.PersonEvent,
	repo EventLocker,
	batchSize int,
	timeout time.Duration,
) *Consumer {
	return &Consumer{
		events:    events,
		repo:      repo,
		batchSize: batchSize,
		timeout:   timeout,
	}
}

func (c *Consumer) Run(ctx context.Context) {
	const op = "Consumer.Run"

	lo.Debug("%s: starting on chan %v", op, c.events)

	for ctx.Err() == nil {

		events, err := c.repo.Lock(uint64(c.batchSize))
		if err != nil {
			sy.Sleep(ctx, c.timeout)
			continue
		}

		lo.Debug("%s: lock events: %v", op, events)

		for _, event := range events {
			if ctx.Err() != nil {
				return
			}

			select {
			case <-ctx.Done():
				return
			case c.events <- &event:
				lo.Debug("%s: sent event: %v", op, event)
			}
		}

		if len(events) < c.batchSize {
			sy.Sleep(ctx, c.timeout)
		}
	}
}
