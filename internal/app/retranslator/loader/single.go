package loader

import (
	"context"
)

type Single struct {
	done chan struct{}
}

func Start(ctx context.Context, r Runner) *Single {
	done := make(chan struct{})
	go func() {
		defer close(done)
		r.Run(ctx)
	}()

	return &Single{
		done: done,
	}
}

func (c *Single) Wait() {
	<-c.done
}
