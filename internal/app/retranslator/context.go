package retranslator

import (
	"context"
)

type contextValueKey int

const termKey contextValueKey = 1

func termFromContext(ctx context.Context) context.Context {
	if term := ctx.Value(termKey).(context.Context); term != nil {
		return term
	}
	return ctx
}

func contextWithTerm(ctx context.Context) (context.Context, context.CancelFunc) {
	term, sendTerm := context.WithCancel(ctx)
	return context.WithValue(ctx, termKey, term), sendTerm
}
