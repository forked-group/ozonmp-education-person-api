package collector

import (
	"context"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/app/retranslator/worker"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/lib/log/lo"
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
	Timeout   time.Duration
}

type collector struct {
	cfg  *Config
	ctx  context.Context
	name string
	buf  []uint64
	n    int
	tm   *time.Timer
}

func (cfg *Config) Run(ctx context.Context, name string) {
	const op = "Run"

	c := collector{
		cfg:  cfg,
		ctx:  ctx,
		name: name,
		buf:  make([]uint64, cfg.BatchSize),
	}

	c.run()
}

func (c *collector) run() {
	const op = "run"

	lo.Debug("%s.%s: running...", c.name, op)

	ok := true

	for ok && c.ctx.Err() == nil {
		ok = c.collect()
		if c.ctx.Err() != nil {
			break
		}
		if !c.flush() {
			return
		}
	}

	if ok {
		lo.Debug("%s.%s: context canceled", c.name, op)
	}
}

func (c *collector) collect() bool {
	const op = "collect"

	if c.tm == nil {
		c.tm = time.NewTimer(c.cfg.Timeout)
	} else {
		c.tm.Reset(c.cfg.Timeout)
	}

loop:
	for c.n < len(c.buf) && c.ctx.Err() == nil {
		select {
		case <-c.ctx.Done():
			break loop

		case <-c.tm.C:
			lo.Debug("%s.%s: collect timeout expired", c.name, op)
			return true

		case event, ok := <-c.cfg.In:
			if !ok {
				lo.Debug("%s.%s: input channel closed", c.name, op)
				if !c.tm.Stop() {
					<-c.tm.C
				}
				return false
			}

			lo.Debug("%s.%s: event %d received", c.name, op, event)
			c.buf[c.n] = event
			c.n++
		}
	}

	if !c.tm.Stop() {
		<-c.tm.C
	}

	if c.n < len(c.buf) {
		lo.Debug("%s.%s: context canceled", c.name, op)
		return false
	}

	lo.Debug("%s.%s: full batch collected", c.name, op)
	return true
}

func (c *collector) flush() bool {
	const op = "flush"

	if c.n == 0 {
		lo.Debug("%s.%s: no any data", c.name, op)
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
		lo.Debug("%s.%s: context canceled", c.name, op)
		return false

	case c.cfg.Out <- job:
		lo.Debug("%s.%s: job %v sent", c.name, op, job)
		return true
	}
}
