package errorhandler

import (
	"context"

	"github.com/lucasanjosmoraes/chicago-ports/pkg/log"
)

// Logger defines what's needed to log information from an error.
type Logger interface {
	Log(context.Context, log.Logger)
	error
}
