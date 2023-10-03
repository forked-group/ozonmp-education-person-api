package retranslator

import (
	"context"
	"time"
)

type consumerConfig struct {
	BatchSize int
	Timeout   time.Duration
	Repo      EventLocker
	Out       chan<- []event
}

type consumer struct {
	cfg  *consumerConfig
	ctx  context.Context
	term context.Context
	tm   *time.Timer
}

func (cfg *consumerConfig) Run(ctx context.Context) {
	c := consumer{
		cfg:  cfg,
		ctx:  ctx,
		term: termFromContext(ctx),
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

func (c *consumer) send(events []event) bool {
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
