package context

import "context"

type contextWithAlert struct {
	context.Context
	alert context.Context
}

func Alert(ctx context.Context) context.Context {
	if c, ok := ctx.(*contextWithAlert); ok {
		return c.alert
	}
	return ctx
}

func WithAlert(ctx context.Context) (*contextWithAlert, context.CancelFunc) {
	alert, sendAlert := context.WithCancel(ctx)
	return &contextWithAlert{
		Context: ctx,
		alert:   alert,
	}, sendAlert
}
