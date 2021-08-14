package http

import (
	"context"
	"io"

	"github.com/lucasanjosmoraes/chicago-ports/pkg/stoppage"
)

// Server defines what's needed to create a HTTP server. It also implements methods
// from the stoppage.Stopper interface to correctly handle its shutdown routine.
type Server interface {
	Listen(ctx context.Context, Router *Router) error
	stoppage.Stopper
}

// HandleFunc defines how an endpoint handler function should be created.
type HandleFunc = func(context.Context, Request, Response)

// Handler defines how a endpoint handler should be created.
type Handler struct {
	Path    string
	Method  string
	Handler HandleFunc
}

// Request defines all the methods needed to manage an HTTP request.
type Request interface {
	Body() io.ReadCloser
	BodyBytes() []byte
	Url() string
	Header(name string) string
	Param(key string) string
	Query(key string) string
}

// Response defines all the methods needed to manage an HTTP response.
type Response interface {
	Write(statuscode int, body []byte) error
	WriteOK(body []byte) error
	WriteCreated(body []byte) error
	WriteNoContent(body []byte) error
	WriteBadRequest(body []byte) error
	WriteUnauthorized(body []byte) error
	WriteForbidden(body []byte) error
	WriteNotfound(body []byte) error
	WriteInternalError(body []byte) error
	JsonResponse()
	XMLResponse()
}
