package loader

import (
	"context"
	"sync"
)

type Group struct {
	done  chan struct{}
	mu    sync.Mutex
	count int
}

func StartGroup(ctx context.Context, runners ...Runner) *Group {
	n := len(runners)

	done := make(chan struct{}, n)

	for _, r := range runners {
		r := r
		go func() {
			defer func() {
				done <- struct{}{}
			}()

			r.Run(ctx)
		}()
	}

	return &Group{
		done:  done,
		count: n,
	}
}

func StartN(ctx context.Context, n int, r Runner) *Group {
	done := make(chan struct{}, n)

	for i := 0; i < n; i++ {
		go func() {
			defer func() {
				done <- struct{}{}
			}()

			r.Run(ctx)
		}()
	}

	return &Group{
		done:  done,
		count: n,
	}
}

func (c *Group) Wait() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for c.count > 0 {
		<-c.done
		c.count--
	}
}
