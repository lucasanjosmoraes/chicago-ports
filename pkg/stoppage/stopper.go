package stoppage

import "context"

// Stopper defines how adapters should handle its shutdown routine.
type Stopper interface {
	Stop(ctx context.Context) error
	StopError() error
}
