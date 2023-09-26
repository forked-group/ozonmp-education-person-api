package worker

import (
	"context"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/lib/log/lo"
	"time"
)

//go:generate mockgen -destination=./mocks/job.go . Job
type Job interface {
	Do() error
}

type Config struct {
	In      <-chan Job
	Timeout time.Duration
}

type worker struct {
	cfg  *Config
	ctx  context.Context
	name string
	tm   *time.Timer
}

func (cfg *Config) Run(ctx context.Context, name string) {
	const op = "Run"

	w := worker{
		cfg:  cfg,
		ctx:  ctx,
		name: name,
	}

	w.run()
}

func (w *worker) run() {
	const op = "run"

	for w.ctx.Err() == nil {
		job, ok := w.receive()
		if !ok {
			return
		}

		err := job.Do()
		if err != nil {
			lo.Debug("%s.%s: can't do job %v: %v", w.name, op, job, err)
			if !w.sleep() {
				return
			}
		}
	}

	lo.Debug("%s.%s: context canceled", w.name, op)
}

func (w *worker) receive() (Job, bool) {
	const op = "receive"

	select {
	case <-w.ctx.Done():
		lo.Debug("%s.%s: context canceled", w.name, op)
		return nil, false

	case job, ok := <-w.cfg.In:
		if !ok {
			lo.Debug("%s.%s: input channel closed", w.name, op)
			return nil, false
		}

		lo.Debug("%s.%s: job %v received", w.name, op, job)
		return job, true
	}
}

func (w *worker) sleep() bool {
	const op = "sleep"

	lo.Debug("%s.%s: sleep for %v", w.name, op, w.cfg.Timeout)

	if w.tm == nil {
		w.tm = time.NewTimer(w.cfg.Timeout)
	} else {
		w.tm.Reset(w.cfg.Timeout)
	}

	select {
	case <-w.ctx.Done():
		lo.Debug("%s.%s: context canceled", w.name, op)
		if !w.tm.Stop() {
			<-w.tm.C
		}
		return false

	case <-w.tm.C:
		lo.Debug("%s.%s: wakeup", w.name, op)
		return true
	}
}
