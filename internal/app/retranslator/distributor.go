package retranslator

import (
	"context"
)

type distributorConfig struct {
	In  <-chan []event
	Out chan<- *event
}

type distributor struct {
	cfg *distributorConfig
	ctx context.Context
}

func (cfg *distributorConfig) Run(ctx context.Context) {
	d := distributor{
		cfg: cfg,
		ctx: ctx,
	}
	d.run()
}

func (d *distributor) run() {
	for d.ctx.Err() == nil {
		events, ok := d.receive()
		if !ok {
			return
		}
		if !d.send(events) {
			return
		}
	}
}

func (d *distributor) receive() ([]event, bool) {
	select {
	case <-d.ctx.Done():
		return nil, false

	case events, ok := <-d.cfg.In:
		if !ok {
			return nil, false
		}
		return events, true
	}
}

func (d *distributor) send(events []event) bool {
	for i := range events {
		event := &events[i]

		select {
		case <-d.ctx.Done():
			return false
		case d.cfg.Out <- event:
		}
	}

	return true
}
