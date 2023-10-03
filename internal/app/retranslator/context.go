package retranslator

import "context"

// TODO: need to use `context.WithValue` instead of type wrapper

type _contextWithTerm struct {
	context.Context
	term context.Context
}

func termFromContext(ctx context.Context) context.Context {
	return ctx.(*_contextWithTerm).term
}

func contextWithTerm(ctx context.Context) (*_contextWithTerm, context.CancelFunc) {
	term, sendTerm := context.WithCancel(ctx)
	return &_contextWithTerm{
		Context: ctx,
		term:    term,
	}, sendTerm
}
