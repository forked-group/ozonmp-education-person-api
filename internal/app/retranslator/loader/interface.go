package loader

import "context"

type Runner interface {
	Run(ctx context.Context, name string)
}
