package retranslator

import (
	"context"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/app/repo"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/app/retranslator/collector"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/app/retranslator/consumer"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/app/retranslator/distributor"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/app/retranslator/loader"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/app/retranslator/producer"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/app/retranslator/worker"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/app/sender"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/model/person"
	"sync"
	"time"
)

type Closer interface {
	CloseCtx(ctx context.Context)
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
	batches chan []person.PersonEvent
	events  chan *person.PersonEvent
	clean   chan uint64
	unlock  chan uint64
	jobs    chan worker.Job

	consumers   Closer
	distributor Closer
	producers   Closer
	cleaner     Closer
	unlocker    Closer
	workers     Closer

	closeTimeout time.Duration
	closeMu      sync.Mutex
	closed       bool
}

func Start(cfg *Config) *Retranslator {
	const op = "Run"

	batches := make(chan []person.PersonEvent)
	events := make(chan *person.PersonEvent, cfg.ChannelSize)
	clean := make(chan uint64)
	unlock := make(chan uint64)
	jobs := make(chan worker.Job)

	consumerCfg := consumer.Config{
		BatchSize: cfg.ConsumeSize,
		Timeout:   cfg.ConsumeTimeout,
		Repo:      cfg.Repo,
		Out:       batches,
	}

	distributorCfg := distributor.Config{
		In:  batches,
		Out: events,
	}

	producerCfg := producer.Config{
		Timeout: cfg.ProduceTimeout,
		In:      events,
		Sender:  cfg.Sender,
		Ok:      clean,
		Error:   unlock,
	}

	cleanerCfg := collector.Config{
		Job:       cfg.Repo.Remove,
		BatchSize: cfg.CollectSize,
		Timeout:   cfg.CollectTimeout,
		In:        clean,
		Out:       jobs,
	}

	unlockerCfg := collector.Config{
		Job:       cfg.Repo.Unlock,
		BatchSize: cfg.CollectSize,
		Timeout:   cfg.CollectTimeout,
		In:        unlock,
		Out:       jobs,
	}

	workerCfg := worker.Config{
		In:      jobs,
		Timeout: cfg.WorkTimeout,
	}

	consumers := loader.StartGroup("Consumer", cfg.ConsumerCount, &consumerCfg)
	distributor := loader.Start("distributor", &distributorCfg)
	producers := loader.StartGroup("Producer", cfg.ProducerCount, &producerCfg)
	cleaner := loader.Start("Cleaner", &cleanerCfg)
	unlocker := loader.Start("Unlocker", &unlockerCfg)
	workers := loader.StartGroup("Worker", cfg.WorkerCount, &workerCfg)

	return &Retranslator{
		batches: batches,
		events:  events,
		clean:   clean,
		unlock:  unlock,
		jobs:    jobs,

		consumers:   consumers,
		distributor: distributor,
		producers:   producers,
		cleaner:     cleaner,
		unlocker:    unlocker,
		workers:     workers,
	}
}

func (r *Retranslator) Close() {
	r.closeMu.Lock()
	defer r.closeMu.Unlock()

	if !r.closed {
		r.closeCtx(context.Background())
	}
}

func (r *Retranslator) CloseCtx(ctx context.Context) {
	r.closeMu.Lock()
	defer r.closeMu.Unlock()

	if !r.closed {
		r.closeCtx(ctx)
	}
}

func (r *Retranslator) closeCtx(ctx context.Context) {
	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		defer wg.Done()
		r.consumers.CloseCtx(ctx)
		close(r.batches)
	}()

	go func() {
		defer wg.Done()
		r.distributor.CloseCtx(ctx)
		close(r.events)
	}()

	go func() {
		defer wg.Done()
		r.producers.CloseCtx(ctx)
		close(r.clean)
		close(r.unlock)
	}()

	go func() {
		defer wg.Done()
		r.cleaner.CloseCtx(ctx)
		r.unlocker.CloseCtx(ctx)
		close(r.jobs)
	}()

	r.workers.CloseCtx(ctx)

	wg.Wait()
	r.closed = true
}
