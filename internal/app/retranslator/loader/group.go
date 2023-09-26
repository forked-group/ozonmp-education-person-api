package loader

import (
	"context"
	"fmt"
	context2 "github.com/aaa2ppp/zonmp-education-person-api/internal/app/retranslator/context"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/lib/log/lo"
	"strings"
	"sync"
)

type Group struct {
	name      string
	cancel    context.CancelFunc
	sendAlert context.CancelFunc
	done      chan struct{}
	closeMu   sync.Mutex
	count     int
}

func StartGroup(name string, n int, r Runner) *Group {
	ctx, cancel := context.WithCancel(context.Background())
	ctx, sendAlert := context2.WithAlert(ctx)
	done := make(chan struct{})

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
		name:      name,
		cancel:    cancel,
		sendAlert: sendAlert,
		done:      done,
		count:     n,
	}
}

func (c *Group) close() {
	const op = "close"

	lo.Debug("%s.%s: cancel context", c.name, op)
	c.cancel()

	// XXX if the runner is poorly written, we can be blocked here
	for c.count > 0 {
		<-c.done
		c.count--
	}

	lo.Debug("%s.%s: all %ss finished", c.name, op, strings.ToLower(c.name))
}

func (c *Group) closeCtx(ctx context.Context) {
	const op = "closeCtx"

	lo.Debug("%s.%s: send alert", c.name, op)
	c.sendAlert()

loop:
	for c.count > 0 {
		select {
		case <-ctx.Done():
			lo.Debug("%s.%s: close context canceled", c.name, op)
			break loop
		case <-c.done:
			c.count--
		}
	}

	if c.count == 0 {
		lo.Debug("%s.%s: all %ss have finished", c.name, op, strings.ToLower(c.name))
		c.cancel() // to release contex resources
		return
	}

	c.close()
}

func (c *Group) Close() {
	c.closeMu.Lock()
	defer c.closeMu.Unlock()

	if c.count == 0 {
		return
	}

	c.sendAlert() // to release alert context resources
	c.close()
}

func (c *Group) CloseCtx(ctx context.Context) {
	c.closeMu.Lock()
	defer c.closeMu.Unlock()

	if c.count == 0 {
		return
	}

	c.closeCtx(ctx)
}
