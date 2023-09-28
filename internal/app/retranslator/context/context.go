package context

import "context"

type contextWithTerm struct {
	context.Context
	term context.Context
}

func Term(ctx context.Context) context.Context {
	return ctx.(*contextWithTerm).term
}

func WithTerm(ctx context.Context) (*contextWithTerm, context.CancelFunc) {
	term, sendTerm := context.WithCancel(ctx)
	return &contextWithTerm{
		Context: ctx,
		term:    term,
	}, sendTerm
}
