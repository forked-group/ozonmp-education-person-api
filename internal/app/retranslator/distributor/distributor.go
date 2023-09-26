package distributor

import (
	"context"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/lib/log/lo"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/model/person"
)

type Config struct {
	In  <-chan []person.PersonEvent
	Out chan<- *person.PersonEvent
}

type distributor struct {
	cfg  *Config
	ctx  context.Context
	name string
}

func (c *Config) Run(ctx context.Context, name string) {
	d := distributor{
		cfg:  c,
		ctx:  ctx,
		name: name,
	}
	d.run()
}

func (d *distributor) run() {
	const op = "run"

	lo.Debug("%s.%s: running...", d.name, op)

	for d.ctx.Err() == nil {
		events, ok := d.receive()
		if !ok {
			return
		}
		if !d.send(events) {
			return
		}
	}

	lo.Debug("%s.%s: context canceled", d.name, op)
}

func (d *distributor) receive() ([]person.PersonEvent, bool) {
	const op = "receive"

	select {
	case <-d.ctx.Done():
		lo.Debug("%s.%s: context canceled", d.name, op)
		return nil, false

	case events, ok := <-d.cfg.In:
		if !ok {
			lo.Debug("%s.%s: input channel closed", d.name, op)
			return nil, false
		}

		lo.Debug("%s.%s: batch received: %v", d.name, op, events)
		return events, true
	}
}

func (d *distributor) send(events []person.PersonEvent) bool {
	const op = "send"

	for i := range events {
		event := &events[i]

		select {
		case <-d.ctx.Done():
			lo.Debug("%s.%s: context canceled", d.name, op)
			return false

		case d.cfg.Out <- event:
			lo.Debug("%s.%s: event %v sent", d.name, op, event)
		}
	}

	return true
}
