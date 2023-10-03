package worker

import (
	"context"
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
	cfg *Config
	ctx context.Context
	tm  *time.Timer
}

func (cfg *Config) Run(ctx context.Context) {
	w := worker{
		cfg: cfg,
		ctx: ctx,
	}
	w.run()
}

func (w *worker) run() {
	for w.ctx.Err() == nil {
		job, ok := w.receive()
		if !ok {
			return
		}

		err := job.Do()
		for err != nil {
			if !w.sleep() {
				return
			}
			err = job.Do()
		}
	}
}

func (w *worker) receive() (Job, bool) {
	select {
	case <-w.ctx.Done():
		return nil, false

	case job, ok := <-w.cfg.In:
		if !ok {
			return nil, false
		}
		return job, true
	}
}

func (w *worker) sleep() bool {
	if w.tm == nil {
		w.tm = time.NewTimer(w.cfg.Timeout)
	} else {
		w.tm.Reset(w.cfg.Timeout)
	}

	select {
	case <-w.ctx.Done():
		if !w.tm.Stop() {
			<-w.tm.C
		}
		return false

	case <-w.tm.C:
		return true
	}
}
