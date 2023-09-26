package sy

import (
	"context"
	"time"
)

func Sleep(ctx context.Context, timeout time.Duration) bool {
	tm := time.NewTimer(timeout)
	select {
	case <-ctx.Done():
		tm.Stop()
		return false
	case <-tm.C:
		return true
	}
}
