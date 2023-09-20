package workerpool

import (
	"context"
	"sync"
	"time"
)

type EventRepo interface {
	Remove(eventIDs []uint64) error
	Unlock(eventIDs []uint64) error
}

type Config struct {
	CleanerCount  int
	UnlockerCount int
	BatchSize     int
	BatchTimeout  time.Duration
	FailTimeout   time.Duration
}

type WorkerPool struct {
	workers     sync.WaitGroup
	batchers    sync.WaitGroup
	clean       chan uint64
	unlock      chan uint64
	batchClean  chan []uint64
	batchUnlock chan []uint64
	cancelBatch context.CancelFunc
	cancelWork  context.CancelFunc
}

func Start(repo EventRepo, cfg Config) *WorkerPool {
	wp := &WorkerPool{
		clean:       make(chan uint64),
		unlock:      make(chan uint64),
		batchClean:  make(chan []uint64),
		batchUnlock: make(chan []uint64),
	}

	wp.start(repo, cfg)
	return wp
}

func (wp *WorkerPool) Close() {
	close(wp.clean)
	close(wp.unlock)
	wp.cancelBatch()
	wp.batchers.Wait()

	close(wp.batchClean)
	close(wp.batchUnlock)
	wp.cancelWork()
	wp.workers.Wait()
}

func (wp *WorkerPool) Remove(ctx context.Context, eventID uint64) {
	select {
	case <-ctx.Done():
	//case wp.batchClean <- []uint64{eventID}: // no batching
	case wp.clean <- eventID:
	}
}

func (wp *WorkerPool) Unlock(ctx context.Context, eventID uint64) {
	select {
	case <-ctx.Done():
	//case wp.batchUnlock <- []uint64{eventID}: // no batching
	case wp.unlock <- eventID:
	}
}

func (wp *WorkerPool) start(repo EventRepo, cfg Config) {
	{
		ctx, cancel := context.WithCancel(context.Background())
		wp.cancelWork = cancel

		wp.startWorkers(cfg.CleanerCount,
			ctx,
			wp.batchClean,
			repo.Remove,
			cfg.FailTimeout,
		)

		wp.startWorkers(cfg.UnlockerCount,
			ctx,
			wp.batchUnlock,
			repo.Unlock,
			cfg.FailTimeout,
		)
	}

	{
		ctx, cancel := context.WithCancel(context.Background())
		wp.cancelBatch = cancel

		wp.startBatcher(ctx, wp.clean, wp.batchClean, cfg)
		wp.startBatcher(ctx, wp.unlock, wp.batchUnlock, cfg)
	}
}

func (wp *WorkerPool) startWorkers(
	n int,
	ctx context.Context,
	ch <-chan []uint64,
	work func(eventIDs []uint64) error,
	timeout time.Duration,
) {
	for i := 0; i < n; i++ {
		wp.workers.Add(1)
		go func() {
			defer wp.workers.Done()
			newWorker(ch, work, timeout).Run(ctx)
		}()
	}
}

func (wp *WorkerPool) startBatcher(ctx context.Context, in <-chan uint64, out chan<- []uint64, cfg Config) {
	wp.batchers.Add(1)
	go func() {
		defer wp.batchers.Done()
		newBatcher(in, out, cfg.BatchSize, cfg.BatchTimeout).Run(ctx)
	}()
}
