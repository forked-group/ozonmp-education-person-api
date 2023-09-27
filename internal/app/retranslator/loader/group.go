package loader

import (
	"context"
	"fmt"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/lib/log/lo"
	"strings"
	"sync"
)

type Group struct {
	name  string
	done  chan struct{}
	mu    sync.Mutex
	count int
}

type Item struct {
	Name string
	Runner
}

func toSingular(name string) string {
	if len(name) > 0 && name[len(name)-1] == 's' {
		return name[:len(name)-1]
	}
	return name
}

func StartGroup(ctx context.Context, name string, group ...Item) *Group {
	name = toSingular(name) // XXX
	n := len(group)

	done := make(chan struct{}, n)

	for _, item := range group {
		item := item
		go func() {
			defer func() {
				lo.Debug("%s: finished", item.Name)
				done <- struct{}{}
			}()

			item.Runner.Run(ctx, item.Name)
		}()
	}

	return &Group{
		name:  name,
		done:  done,
		count: n,
	}
}

func StartN(ctx context.Context, name string, n int, r Runner) *Group {
	name = toSingular(name) // XXX

	done := make(chan struct{}, n)

	for i := 0; i < n; i++ {
		itemName := fmt.Sprintf("%s%d", name, i+1)

		go func() {
			defer func() {
				lo.Debug("%s: finished", itemName)
				done <- struct{}{}
			}()

			r.Run(ctx, itemName)
		}()
	}

	return &Group{
		name:  name,
		done:  done,
		count: n,
	}
}

func (c *Group) Wait() {
	const op = "Wait"

	c.mu.Lock()
	defer c.mu.Unlock()

	for c.count > 0 {
		<-c.done
		c.count--
	}

	lo.Debug("%s.%s: all %ss finished", c.name, op, strings.ToLower(c.name))
}
