package database

import (
	"context"
	"fmt"

	"github.com/lucasanjosmoraes/chicago-ports/pkg/log"
)

// NoRowsAffectedError can be used in execution statements (like insert, update, delete, etc.) that returns how many rows
// have been affected.
type NoRowsAffectedError struct {
	Stmt      string
	NeedRetry bool
	SendToDLQ bool
}

// Log implements errorhandler.Logger.
func (e NoRowsAffectedError) Log(ctx context.Context, l log.Logger) {
	l.Errorf(ctx, e.buildMessage())
}

// Error implements error interface.
func (e NoRowsAffectedError) Error() string {
	return e.buildMessage()
}

// Retryable implements subscriber.Error.
func (e NoRowsAffectedError) Retryable() bool {
	return e.NeedRetry
}

// DLSendable implements subscriber.Error.
func (e NoRowsAffectedError) DLSendable() bool {
	return e.SendToDLQ
}

func (e NoRowsAffectedError) buildMessage() string {
	return fmt.Sprintf("no rows affected executing '%s'", e.Stmt)
}
