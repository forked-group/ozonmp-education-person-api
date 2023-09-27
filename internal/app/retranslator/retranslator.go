package retranslator

import (
	"context"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/app/repo"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/app/retranslator/collector"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/app/retranslator/consumer"
	context2 "github.com/aaa2ppp/zonmp-education-person-api/internal/app/retranslator/context"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/app/retranslator/distributor"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/app/retranslator/loader"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/app/retranslator/producer"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/app/retranslator/worker"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/app/sender"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/model/person"
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

	CollectSize    int
	CollectTimeout time.Duration

	WorkerCount int
	WorkTimeout time.Duration

	Repo   repo.EventRepo
	Sender sender.EventSender
}

type Retranslator struct {
	sendTerm context.CancelFunc

	batches chan []person.PersonEvent
	events  chan *person.PersonEvent
	clean   chan uint64
	unlock  chan uint64
	jobs    chan worker.Job

	consumers   Waiter
	distributor Waiter
	producers   Waiter
	collectors  Waiter
	workers     Waiter

	mu     sync.Mutex
	closed bool
}

func (cfg *Config) Start(ctx context.Context) *Retranslator {
	batches := make(chan []person.PersonEvent)
	events := make(chan *person.PersonEvent, cfg.ChannelSize)
	clean := make(chan uint64)
	unlock := make(chan uint64)
	jobs := make(chan worker.Job)

	termCtx, sendTerm := context2.WithTerm(ctx)

	consumers := loader.StartN(termCtx, "Consumers", cfg.ConsumerCount,
		&consumer.Config{
			BatchSize: cfg.ConsumeSize,
			Timeout:   cfg.ConsumeTimeout,
			Repo:      cfg.Repo,
			Out:       batches,
		},
	)

	distributor := loader.Start(ctx, "Distributor",
		&distributor.Config{
			In:  batches,
			Out: events,
		},
	)

	producers := loader.StartN(ctx, "Producers", cfg.ProducerCount,
		&producer.Config{
			Timeout: cfg.ProduceTimeout,
			In:      events,
			Sender:  cfg.Sender,
			Ok:      clean,
			Error:   unlock,
		},
	)

	collectors := loader.StartGroup(ctx, "Collectors",
		loader.Item{"Cleaner", &collector.Config{
			Job:       cfg.Repo.Remove,
			BatchSize: cfg.CollectSize,
			Timeout:   cfg.CollectTimeout,
			In:        clean,
			Out:       jobs,
		}},
		loader.Item{"Unlocker", &collector.Config{
			Job:       cfg.Repo.Unlock,
			BatchSize: cfg.CollectSize,
			Timeout:   cfg.CollectTimeout,
			In:        unlock,
			Out:       jobs,
		}},
	)

	workers := loader.StartN(ctx, "Workers", cfg.WorkerCount,
		&worker.Config{
			In:      jobs,
			Timeout: cfg.WorkTimeout,
		},
	)

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
