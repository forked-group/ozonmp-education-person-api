package loader

import (
	"context"
	context2 "github.com/aaa2ppp/zonmp-education-person-api/internal/app/retranslator/context"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/lib/log/lo"
	"sync"
)

type Single struct {
	name      string
	cancel    context.CancelFunc
	sendAlert context.CancelFunc
	done      chan struct{}
	once      sync.Once
	closeMu   sync.Mutex
	closed    bool
}

func Start(name string, r Runner) *Single {
	ctx, cancel := context.WithCancel(context.Background())
	ctx, sendAlert := context2.WithAlert(ctx)

	done := make(chan struct{})
	go func() {
		defer close(done)
		r.Run(ctx, name)
	}()

	return &Single{
		name:      name,
		cancel:    cancel,
		sendAlert: sendAlert,
		done:      done,
	}
}

func (c *Single) close() {
	const op = "close"

	lo.Debug("%s.%s: cancel context", c.name, op)
	c.cancel()

	// XXX if the runner is poorly written, we can be blocked here
	<-c.done
	lo.Debug("%s.%s: finished", c.name, op)

	c.closed = true
}

func (c *Single) closeCtx(ctx context.Context) {
	const op = "closeCtx"

	lo.Debug("%s.%s: send alert", c.name, op)
	c.sendAlert()

	select {
	case <-ctx.Done():
		lo.Debug("%s.%s: close context canceled", c.name, op)
		c.close()

	case <-c.done:
		c.cancel() // to release context resources
		lo.Debug("%s.%s: finished", c.name, op)
	}

	c.closed = true
}

func (c *Single) Close() {
	c.closeMu.Lock()
	defer c.closeMu.Unlock()

	if c.closed {
		return
	}

	c.sendAlert() // to release alert context resources
	c.close()
}

func (c *Single) CloseCtx(ctx context.Context) {
	c.closeMu.Lock()
	defer c.closeMu.Unlock()

	if c.closed {
		return
	}

	c.closeCtx(ctx)
}
