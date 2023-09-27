package loader

import (
	"context"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/lib/log/lo"
)

type Single struct {
	name string
	done chan struct{}
}

func Start(ctx context.Context, name string, r Runner) *Single {
	done := make(chan struct{})
	go func() {
		defer close(done)
		r.Run(ctx, name)
	}()

	return &Single{
		name: name,
		done: done,
	}
}

func (c *Single) Wait() {
	const op = "Wait"

	<-c.done
	lo.Debug("%s.%s: finished", c.name, op)
}
