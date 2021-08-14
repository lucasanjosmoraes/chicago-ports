package database

import (
	"context"

	"github.com/lucasanjosmoraes/chicago-ports/pkg/log"
)

// Error can be used when you get errors from data sources that connects with databases.
type Error struct {
	Stmt      string
	Err       error
	NeedRetry bool
	SendToDLQ bool
}

// Log implements errorhandler.Logger.
func (e Error) Log(ctx context.Context, l log.Logger) {
	l.Errorf(ctx, "error executing database statement '%s': %s", e.Stmt, e.Err)
}

// Error implements error interface.
func (e Error) Error() string {
	return e.Err.Error()
}

// Retryable implements subscriber.Error.
func (e Error) Retryable() bool {
	return e.NeedRetry
}

// DLSendable implements subscriber.Error.
func (e Error) DLSendable() bool {
	return e.SendToDLQ
}
