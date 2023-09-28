package collector

import (
	"context"
	"github.com/aaa2ppp/ozonmp-education-kw-person-api/internal/app/retranslator/worker"
	"time"
)

type batchJob struct {
	do    func([]uint64) error
	batch []uint64
}

func (j batchJob) Do() error {
	return j.do(j.batch)
}

type Config struct {
	Job       func([]uint64) error
	In        <-chan uint64
	Out       chan<- worker.Job
	BatchSize int
	MaxDelay  time.Duration
}

type collector struct {
	cfg *Config
	ctx context.Context
	buf []uint64
	n   int
	tm  *time.Timer
}

func (cfg *Config) Run(ctx context.Context) {
	c := collector{
		cfg: cfg,
		ctx: ctx,
		buf: make([]uint64, cfg.BatchSize),
	}
	c.run()
}

func (c *collector) run() {
	for ok := true; ok && c.ctx.Err() == nil; {
		ok = c.collect()
		if c.ctx.Err() != nil {
			break
		}

		if !c.flush() {
			return
		}
	}
}

func (c *collector) setTimer(delay time.Duration) {
	if c.tm == nil {
		c.tm = time.NewTimer(delay)
	} else {
		c.tm.Reset(delay)
	}
}

func (c *collector) collect() bool {
	select {
	case <-c.ctx.Done():
		return false

	case event, ok := <-c.cfg.In:
		if !ok {
			return false
		}

		c.setTimer(c.cfg.MaxDelay)
		c.buf[0] = event
		c.n = 1
	}

loop:
	for c.n < len(c.buf) && c.ctx.Err() == nil {
		select {
		case <-c.ctx.Done():
			break loop

		case <-c.tm.C:
			return true

		case event, ok := <-c.cfg.In:
			if !ok {
				if !c.tm.Stop() {
					<-c.tm.C
				}
				return false
			}

			c.buf[c.n] = event
			c.n++
		}
	}

	if !c.tm.Stop() {
		<-c.tm.C
	}

	return c.n == len(c.buf)
}

func (c *collector) flush() bool {
	if c.n == 0 {
		return true
	}

	batch := make([]uint64, c.n)
	copy(batch, c.buf)
	c.n = 0

	job := batchJob{
		do:    c.cfg.Job,
		batch: batch,
	}

	select {
	case <-c.ctx.Done():
		return false

	case c.cfg.Out <- job:
		return true
	}
}
