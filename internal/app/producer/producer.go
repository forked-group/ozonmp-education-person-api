package producer

import (
	"context"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/app/workerpool"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/lib/log/lo"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/lib/sync/sy"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/model/person"
	"time"
)

type EventSender interface {
	Send(person *person.PersonEvent) error
}

type Producer struct {
	events     chan *person.PersonEvent
	sender     EventSender
	workerPoll *workerpool.WorkerPool
	timeout    time.Duration
}

func New(
	events chan *person.PersonEvent,
	sender EventSender,
	workerPool *workerpool.WorkerPool,
	timeout time.Duration,
) *Producer {
	return &Producer{
		events:     events,
		sender:     sender,
		workerPoll: workerPool,
		timeout:    timeout,
	}
}

func (p *Producer) Run(ctx context.Context) {
	const op = "Producer.Run"

	lo.Debug("%s: starting on chan %v", op, p.events)

loop:
	for ctx.Err() == nil {

		select {
		case <-ctx.Done():
			return

		case event, ok := <-p.events:
			if !ok {
				return
			}

			lo.Debug("%s: got event: %v", op, event)

			if err := p.sender.Send(event); err != nil {
				p.workerPoll.Unlock(ctx, event.ID)
				sy.Sleep(ctx, p.timeout)
				continue loop
			}

			p.workerPoll.Remove(ctx, event.ID)
		}

	}
}
