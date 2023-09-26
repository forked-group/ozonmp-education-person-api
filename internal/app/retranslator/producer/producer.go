package producer

import (
	"context"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/lib/log/lo"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/model/person"
	"time"
)

//go:generate mockgen -destination=./mocks/event_sender.go . EventSender
type EventSender interface {
	Send(person *person.PersonEvent) error
}

type Config struct {
	Timeout time.Duration
	In      <-chan *person.PersonEvent
	Sender  EventSender
	Ok      chan<- uint64
	Error   chan<- uint64
}

type producer struct {
	cfg  *Config
	ctx  context.Context
	name string
	tm   *time.Timer
}

func (cfg *Config) Run(ctx context.Context, name string) {
	const op = "Run"

	p := producer{
		cfg:  cfg,
		ctx:  ctx,
		name: name,
	}

	p.run()
}

func (p *producer) run() {
	const op = "run"

	lo.Debug("%s.%s: running...", p.name, op)

	for p.ctx.Err() == nil {
		event, ok := p.receive()
		if !ok {
			return
		}

		if err := p.cfg.Sender.Send(event); err != nil {
			lo.Debug("%s.%s: can't send event %v: %v", p.name, op, event, err)
			if !p.report(p.cfg.Error, event.ID) || !p.sleep() {
				return
			}
		} else {
			if !p.report(p.cfg.Ok, event.ID) {
				return
			}
		}
	}

	lo.Debug("%s.%s: context canceled", p.name, op)
}

func (p *producer) receive() (*person.PersonEvent, bool) {
	const op = "receive"

	select {
	case <-p.ctx.Done():
		lo.Debug("%s.%s: context canceled", p.name, op)
		return nil, false

	case event, ok := <-p.cfg.In:
		if !ok {
			lo.Debug("%s.%s: input channel closed", p.name, op)
			return nil, false
		}

		lo.Debug("%s.%s: event %v received", p.name, op, event)
		return event, true
	}
}

func (p *producer) report(out chan<- uint64, eventID uint64) bool {
	const op = "report"

	select {
	case <-p.ctx.Done():
		lo.Debug("%s.%s: context canceled", p.name, op)
		return false

	case out <- eventID:
		lo.Debug("%s.%s: report sent", p.name, op)
		return true
	}
}

func (p *producer) sleep() bool {
	const op = "sleep"

	lo.Debug("%s.%s: sleep for %v", p.name, op, p.cfg.Timeout)

	if p.tm == nil {
		p.tm = time.NewTimer(p.cfg.Timeout)
	} else {
		p.tm.Reset(p.cfg.Timeout)
	}

	select {
	case <-p.ctx.Done():
		lo.Debug("%s.%s: context canceled", p.name, op)
		if !p.tm.Stop() {
			<-p.tm.C
		}
		return false

	case <-p.tm.C:
		lo.Debug("%s.%s: wakeup", p.name, op)
		return true
	}
}
