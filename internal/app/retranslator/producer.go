package retranslator

import (
	"context"
	"time"
)

type producerConfig struct {
	Timeout time.Duration
	In      <-chan *event
	Sender  EventSender
	Ok      chan<- uint64
	Error   chan<- uint64
}

type producer struct {
	cfg *producerConfig
	ctx context.Context
	tm  *time.Timer
}

func (cfg *producerConfig) Run(ctx context.Context) {
	p := producer{
		cfg: cfg,
		ctx: ctx,
	}
	p.run()
}

func (p *producer) run() {
	for p.ctx.Err() == nil {
		event, ok := p.receive()
		if !ok {
			return
		}

		if err := p.cfg.Sender.Send(event); err != nil {
			if !p.report(p.cfg.Error, event.ID) || !p.sleep() {
				return
			}
		} else {
			if !p.report(p.cfg.Ok, event.ID) {
				return
			}
		}
	}
}

func (p *producer) receive() (*event, bool) {
	select {
	case <-p.ctx.Done():
		return nil, false

	case event, ok := <-p.cfg.In:
		if !ok {
			return nil, false
		}
		return event, true
	}
}

func (p *producer) report(out chan<- uint64, eventID uint64) bool {
	select {
	case <-p.ctx.Done():
		return false

	case out <- eventID:
		return true
	}
}

func (p *producer) sleep() bool {
	if p.tm == nil {
		p.tm = time.NewTimer(p.cfg.Timeout)
	} else {
		p.tm.Reset(p.cfg.Timeout)
	}

	select {
	case <-p.ctx.Done():
		if !p.tm.Stop() {
			<-p.tm.C
		}
		return false

	case <-p.tm.C:
		return true
	}
}
