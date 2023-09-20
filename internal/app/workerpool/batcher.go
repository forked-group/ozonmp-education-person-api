package workerpool

import (
	"context"
	"fmt"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/lib/log/lo"
	"time"
)

type batcher struct {
	ctx       context.Context
	in        <-chan uint64
	out       chan<- []uint64
	batchSize int
	batch     []uint64
	timeout   time.Duration
}

func newBatcher(
	in <-chan uint64,
	out chan<- []uint64,
	batchSize int,
	timeout time.Duration,
) *batcher {
	if batchSize < 0 {
		panic(fmt.Sprintf("newBatcher: got negative batchSize %d", batchSize))
	}

	return &batcher{
		in:        in,
		out:       out,
		batchSize: batchSize,
		timeout:   timeout,
	}
}

func (b *batcher) Run(ctx context.Context) {
	const op = "batcher.Run"

	lo.Debug("%s: starting on in chan %v", op, b.in)

	b.batch = make([]uint64, 0, b.batchSize)
	t := time.NewTicker(b.timeout)

loop:
	for ctx.Err() == nil {
		select {
		case <-ctx.Done():
			return

		case id, ok := <-b.in:
			if !ok {
				break loop
			}

			lo.Debug("%s: %v: got id %d", op, b.in, id)

			b.batch = append(b.batch, id)

			if len(b.batch) >= b.batchSize {
				lo.Debug("%s: %v: batch completed", op, b.in)

				b.flush(ctx)
				b.batch = make([]uint64, 0, b.batchSize)
				t.Reset(b.timeout)
			}

		case <-t.C:
			lo.Debug("%s: %v: timeout expired", op, b.in)

			if len(b.batch) != 0 {
				b.flush(ctx)
				b.batch = make([]uint64, 0, b.batchSize)
			}
		}
	}

	if len(b.batch) != 0 {
		b.flush(ctx)
	}
}

func (b *batcher) flush(ctx context.Context) {
	const op = "batcher.flush"

	completed := b.batch
	b.batch = nil

	select {
	case <-ctx.Done():
	case b.out <- completed:
		lo.Debug("%s: %v: sent batch: %v", op, b.in, completed)
	}
}
