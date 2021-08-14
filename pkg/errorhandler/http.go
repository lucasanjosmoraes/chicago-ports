package errorhandler

import (
	"context"

	"github.com/lucasanjosmoraes/chicago-ports/pkg/http"
	"github.com/lucasanjosmoraes/chicago-ports/pkg/log"
)

// HTTPHandler defines what's needed to handle errors returned from a http.Response
type HTTPHandler interface {
	Handle(context.Context, http.Response, error)
}

// ResponseWithLog implements HTTPHandler.
type ResponseWithLog struct {
	Logger log.Logger
}

// Handle validates if the given error implements Logger or Responder to call its methods.
func (h ResponseWithLog) Handle(ctx context.Context, res http.Response, err error) {
	h.log(ctx, err)

	appErr, ok := err.(Responder)
	if ok && appErr != nil {
		writeErr := res.Write(appErr.Status(), appErr.Response())
		if writeErr != nil {
			h.Logger.Errorf(ctx, "error writing response: %s", writeErr)
		}
		return
	}

	writeErr := res.WriteInternalError([]byte(""))
	if writeErr != nil {
		h.Logger.Errorf(ctx, "error writing internal error response: %s", writeErr)
	}
}

func (h ResponseWithLog) log(ctx context.Context, err error) {
	if err == nil {
		return
	}

	logErr, ok := err.(Logger)
	if ok && logErr != nil {
		logErr.Log(ctx, h.Logger)
		return
	}

	h.Logger.Errorf(ctx, err.Error())
}
