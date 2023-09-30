package consumer

import (
	"context"
	context2 "github.com/aaa2ppp/ozonmp-education-person-api/internal/app/retranslator/context"
	"time"
)

//go:generate mockgen -destination=./mocks/event_locker.go . EventLocker
type EventLocker interface {
	Lock(n uint64) ([]education.PersonEvent, error)
}

type Config struct {
	BatchSize int
	Timeout   time.Duration
	Repo      EventLocker
	Out       chan<- []education.PersonEvent
}

type consumer struct {
	cfg  *Config
	ctx  context.Context
	term context.Context
	tm   *time.Timer
}

func (cfg *Config) Run(ctx context.Context) {
	c := consumer{
		cfg:  cfg,
		ctx:  ctx,
		term: context2.Term(ctx),
	}
	c.run()
}

func (c *consumer) run() {
	for c.term.Err() == nil {
		events, err := c.cfg.Repo.Lock(uint64(c.cfg.BatchSize))
		if err != nil {
			if !c.sleep() {
				return
			}
			continue
		}

		if !c.send(events) {
			return
		}

		if len(events) < c.cfg.BatchSize {
			if !c.sleep() {
				return
			}
		}
	}
}

func (c *consumer) send(events []education.PersonEvent) bool {
	select {
	case <-c.ctx.Done():
		return false

	case c.cfg.Out <- events:
		return true
	}
}

func (c *consumer) sleep() bool {
	if c.tm == nil {
		c.tm = time.NewTimer(c.cfg.Timeout)
	} else {
		c.tm.Reset(c.cfg.Timeout)
	}

	select {
	case <-c.term.Done():
		if !c.tm.Stop() {
			<-c.tm.C
		}
		return false

	case <-c.tm.C:
		return true
	}
}
