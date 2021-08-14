package stoppage

import (
	"context"
	"os"
	"os/signal"
	"syscall"
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
	ctx, done := context.WithCancel(ctx)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		defer done()

		signalCalled := <-quit
		s.Logger.Infof(ctx, "Starting shutdown by signal: ", signalCalled.String())
		signal.Stop(quit)
		close(quit)

		s.graceful(ctx)
	}()

	return ctx
}

func (s Shutdown) graceful(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	for _, stopper := range s.stoppers {
		err := stopper.Stop(ctx)
		if err != nil {
			s.Logger.Error(ctx, err.Error())
		}
	}
}
