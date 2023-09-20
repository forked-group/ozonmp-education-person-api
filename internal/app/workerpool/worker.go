package workerpool

import (
	"context"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/lib/log/lo"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/lib/sync/sy"
	"time"
)

type worker struct {
	ch      <-chan []uint64
	work    func(eventIDs []uint64) error
	timeout time.Duration
}

func newWorker(
	ch <-chan []uint64,
	work func(eventIDs []uint64) error,
	timeout time.Duration,
) *worker {
	return &worker{
		ch:      ch,
		work:    work,
		timeout: timeout,
	}
}

func (w *worker) Run(ctx context.Context) {
	const op = "worker.Run"

	for ctx.Err() == nil {
		select {
		case <-ctx.Done():
			return

		case batchIDs, ok := <-w.ch:
			if !ok {
				return
			}

			lo.Debug("%s: got batch: %v", op, batchIDs)

			if err := w.work(batchIDs); err != nil {
				lo.Debug("%s: can't do work: %v", op, err)
				sy.Sleep(ctx, w.timeout)
			}
		}
	}
}
