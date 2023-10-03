package retranslator

import (
	"context"
	"sync"
	"time"
)

type Waiter interface {
	Wait()
}

type Config struct {
	ChannelSize int

	ConsumerCount  int
	ConsumeSize    int
	ConsumeTimeout time.Duration

	ProducerCount  int
	ProduceTimeout time.Duration

	CollectSize     int
	CollectMaxDelay time.Duration

	WorkerCount int
	WorkTimeout time.Duration

	Repo   EventRepo
	Sender EventSender
}

type Retranslator struct {
	sendTerm context.CancelFunc

	batches chan []event
	events  chan *event
	clean   chan uint64
	unlock  chan uint64
	jobs    chan workerJob

	consumers   Waiter
	distributor Waiter
	producers   Waiter
	collectors  Waiter
	workers     Waiter

	mu     sync.Mutex
	closed bool
}

func (cfg *Config) Start(ctx context.Context) *Retranslator {
	batches := make(chan []event)
	events := make(chan *event, cfg.ChannelSize)
	clean := make(chan uint64)
	unlock := make(chan uint64)
	jobs := make(chan workerJob)

	termCtx, sendTerm := contextWithTerm(ctx)

	consumers := startN(termCtx, cfg.ConsumerCount,
		&consumerConfig{
			BatchSize: cfg.ConsumeSize,
			Timeout:   cfg.ConsumeTimeout,
			Repo:      cfg.Repo,
			Out:       batches,
		})

	distributor := start(ctx,
		&distributorConfig{
			In:  batches,
			Out: events,
		})

	producers := startN(ctx, cfg.ProducerCount,
		&producerConfig{
			Timeout: cfg.ProduceTimeout,
			In:      events,
			Sender:  cfg.Sender,
			Ok:      clean,
			Error:   unlock,
		})

	collectors := groupStart(ctx,
		&collectorConfig{
			Job:       cfg.Repo.Remove,
			BatchSize: cfg.CollectSize,
			MaxDelay:  cfg.CollectMaxDelay,
			In:        clean,
			Out:       jobs,
		},
		&collectorConfig{
			Job:       cfg.Repo.Unlock,
			BatchSize: cfg.CollectSize,
			MaxDelay:  cfg.CollectMaxDelay,
			In:        unlock,
			Out:       jobs,
		})

	workers := startN(ctx, cfg.WorkerCount,
		&workerConfig{
			In:      jobs,
			Timeout: cfg.WorkTimeout,
		})

	return &Retranslator{
		sendTerm: sendTerm,

		batches: batches,
		events:  events,
		clean:   clean,
		unlock:  unlock,
		jobs:    jobs,

		consumers:   consumers,
		distributor: distributor,
		producers:   producers,
		collectors:  collectors,
		workers:     workers,
	}
}

func (r *Retranslator) Close() {
	r.mu.Lock()
	defer r.mu.Unlock()

	if !r.closed {
		r.close()
		r.closed = true
	}
}

func (r *Retranslator) close() {
	r.sendTerm()

	r.consumers.Wait()
	close(r.batches)

	r.distributor.Wait()
	close(r.events)

	r.producers.Wait()
	close(r.clean)
	close(r.unlock)

	r.collectors.Wait()
	close(r.jobs)

	r.workers.Wait()
}
