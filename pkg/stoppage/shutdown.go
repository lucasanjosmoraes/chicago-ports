package stoppage

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/lucasanjosmoraes/chicago-ports/pkg/log"
)

// Shutdown helps to manage many stoppers when we need to shutdown an application.
type Shutdown struct {
	Logger   log.Logger
	stoppers []Stopper
}

// NewShutdown instantiates a new Shutdown.
func NewShutdown(logger log.Logger) Shutdown {
	return Shutdown{
		Logger:   logger,
		stoppers: nil,
	}
}

// Add will add a new stopper to the stoppers list of the shutdown pointer.
func (s *Shutdown) Add(stopper Stopper) {
	s.stoppers = append(s.stoppers, stopper)
}

// GracefulSignal accepts a context and it returns a new one that handles signals
// from the system to shutdown the application.
func (s Shutdown) GracefulSignal(ctx context.Context) context.Context {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	done := make(chan bool, 1)
	stop := make(chan bool, 1)

	go func() {
		for {
			select {
			case <-ctx.Done():
				stop <- true
			case <-stop:
				s.graceful(ctx, done)
				cancel()
				return
			}
		}
	}()

	return ctx
}

func (s Shutdown) graceful(ctx context.Context, done chan<- bool) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	for _, stopper := range s.stoppers {
		err := stopper.Stop(ctx)
		if err != nil {
			s.Logger.Error(ctx, err.Error())
		}
	}

	done <- true
}
