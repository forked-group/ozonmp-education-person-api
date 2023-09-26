package consumer

import (
	"context"
	context2 "github.com/aaa2ppp/zonmp-education-person-api/internal/app/retranslator/context"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/lib/log/lo"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/model/person"
	"time"
)

//go:generate mockgen -destination=./mocks/event_locker.go . EventLocker
type EventLocker interface {
	Lock(n uint64) ([]person.PersonEvent, error)
}

type Config struct {
	BatchSize int
	Timeout   time.Duration
	Repo      EventLocker
	Out       chan<- []person.PersonEvent
}

type consumer struct {
	cfg   *Config
	name  string
	ctx   context.Context
	alert context.Context
	tm    *time.Timer
}

func (cfg *Config) Run(ctx context.Context, name string) {
	const op = "Run"

	c := consumer{
		cfg:   cfg,
		name:  name,
		ctx:   ctx,
		alert: context2.Alert(ctx),
	}

	c.run()
}

func (c *consumer) run() {
	const op = "run"

	lo.Debug("%s.%s: running...", c.name, op)

	for c.alert.Err() == nil {
		events, err := c.cfg.Repo.Lock(uint64(c.cfg.BatchSize))
		if err != nil {
			lo.Debug("%s.%s: can't lock events: %v", c.name, op, err)
			if !c.sleep() {
				return
			}
			continue
		}

		lo.Debug("%s.%s: events locked: %v", c.name, op, events)

		if !c.send(events) {
			return
		}

		if len(events) < c.cfg.BatchSize {
			lo.Debug("%s.%s: incomplete batch (%d/%d) was received",
				c.name, op, len(events), c.cfg.BatchSize)
			if !c.sleep() {
				return
			}
		}
	}

	lo.Debug("%s.%s: alert received", c.name, op)
}

func (c *consumer) send(events []person.PersonEvent) bool {
	const op = "send"

	select {
	case <-c.ctx.Done():
		lo.Debug("%s.%s: context canceled", c.name, op)
		return false

	case c.cfg.Out <- events:
		lo.Debug("%s.%s: event batch sent: %v", c.name, op, events)
	}

	lo.Debug("%s.%s: all events sent", c.name, op)
	return true
}

func (c *consumer) sleep() bool {
	const op = "sleep"

	lo.Debug("%s.%s: sleep for %v", c.name, op, c.cfg.Timeout)

	if c.tm == nil {
		c.tm = time.NewTimer(c.cfg.Timeout)
	} else {
		c.tm.Reset(c.cfg.Timeout)
	}

	select {
	case <-c.alert.Done():
		lo.Debug("%s.%s: alert received", c.name, op)
		if !c.tm.Stop() {
			<-c.tm.C
		}
		return false

	case <-c.tm.C:
		lo.Debug("%s.%s: wakeup", c.name, op)
		return true
	}
}
