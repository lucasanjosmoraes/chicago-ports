package subscriber

import (
	"context"

	"github.com/lucasanjosmoraes/chicago-ports/pkg/log"
)

// Logger defines what's needed to log information from an error.
type Logger interface {
	Log(context.Context, log.Logger)
}

// Error defines how it should be handled by event consumers.
type Error interface {
	Retryable() bool
	DLSendable() bool
}

// DLBehavior allows devs to customize sending messages to DL.
type DLBehavior interface {
	DLErrorMessage() string
}

// Log accepts an error and will call its Log method, if it implements Logger.
// Otherwise will log the return of its Error method.
func Log(ctx context.Context, err error, logger log.Logger) {
	if err == nil {
		return
	}

	logErr, ok := err.(Logger)
	if ok {
		logErr.Log(ctx, logger)
		return
	}

	logger.Errorf(ctx, err.Error())
}
