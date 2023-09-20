package sy

import (
	"context"
	"time"
)

func Sleep(ctx context.Context, timeout time.Duration) {
	t := time.NewTimer(timeout)
	select {
	case <-ctx.Done():
	case <-t.C:
	}
}
