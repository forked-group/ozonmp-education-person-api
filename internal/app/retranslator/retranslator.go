package retranslator

import (
	"context"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/app/consumer"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/app/producer"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/app/repo"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/app/sender"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/app/workerpool"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/model/person"
	"sync"
	"time"
)

type Config struct {
	ChannelSize int

	ConsumerCount     int
	ConsumerBatchSize int
	ConsumerTimeout   time.Duration

	ProducerCount   int
	ProducerTimeout time.Duration

	WorkerPoolCfg workerpool.Config
}

type Retranslator struct {
	events        chan *person.PersonEvent
	consumers     sync.WaitGroup
	producers     sync.WaitGroup
	cancelConsume context.CancelFunc
	cancelProduce context.CancelFunc
	workerPool    *workerpool.WorkerPool
	closeMu       sync.Mutex
	closed        bool
}

func Start(repo repo.EventRepo, sender sender.EventSender, cfg Config) *Retranslator {
	r := &Retranslator{
		events:     make(chan *person.PersonEvent, cfg.ChannelSize),
		workerPool: workerpool.Start(repo, cfg.WorkerPoolCfg),
	}

	r.startProduce(sender, cfg)
	r.startConsume(repo, cfg)

	return r
}

func (r *Retranslator) startProduce(sender sender.EventSender, cfg Config) {
	ctx, cancel := context.WithCancel(context.Background())
	r.cancelProduce = cancel

	for i := 0; i < cfg.ProducerCount; i++ {
		r.producers.Add(1)
		go func() {
			defer r.producers.Done()
			producer.New(
				r.events,
				sender,
				r.workerPool,
				cfg.ProducerTimeout,
			).Run(ctx)
		}()
	}
}

func (r *Retranslator) startConsume(repo repo.EventRepo, cfg Config) {
	ctx, cancel := context.WithCancel(context.Background())
	r.cancelConsume = cancel

	for i := 0; i < cfg.ConsumerCount; i++ {
		r.consumers.Add(1)
		go func() {
			defer r.consumers.Done()
			consumer.New(
				r.events,
				repo,
				cfg.ConsumerBatchSize,
				cfg.ConsumerTimeout,
			).Run(ctx)
		}()
	}
}

func (r *Retranslator) Close() {
	r.closeMu.Lock()
	defer r.closeMu.Unlock()

	if r.closed {
		return
	}

	r.cancelConsume()
	r.consumers.Wait()

	close(r.events)
	r.cancelProduce()
	r.producers.Wait()

	r.workerPool.Close()

	r.closed = true
}
