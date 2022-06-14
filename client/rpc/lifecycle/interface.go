package lifecycle

import "context"

type Lifecycler interface {
	Heartbeat(ctx context.Context) error
	UnRegistry(ctx context.Context) error
}
