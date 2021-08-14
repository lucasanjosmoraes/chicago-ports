package http

import "context"

// StaticContent contains all attributes needed to serve endpoints from static files.
type StaticContent struct {
	Path string
	Root string
}

// Router defines which HTTP verbs are supported by the Server.
type Router struct {
	Handlers       []Handler
	StaticHandlers []StaticContent
}

func NewRouter() *Router {
	return new(Router)
}

func (r *Router) Get(url string, handler HandleFunc) {
	r.Handlers = createHandler(r.Handlers, url, "GET", handler)
}

func (r *Router) Post(url string, handler HandleFunc) {
	r.Handlers = createHandler(r.Handlers, url, "POST", handler)
}

func (r *Router) Put(url string, handler HandleFunc) {
	r.Handlers = createHandler(r.Handlers, url, "PUT", handler)
}

func (r *Router) Patch(url string, handler HandleFunc) {
	r.Handlers = createHandler(r.Handlers, url, "PATCH", handler)
}

func (r *Router) Delete(url string, handler HandleFunc) {
	r.Handlers = createHandler(r.Handlers, url, "DELETE", handler)
}

func (r *Router) AddStaticContent(path, root string) {
	r.StaticHandlers = createStaticHandler(r.StaticHandlers, path, root)
}

func createHandler(handlers []Handler, url string, method string, handler func(context.Context, Request, Response)) []Handler {
	return append(handlers, Handler{
		Path:    url,
		Method:  method,
		Handler: handler,
	})
}

func createStaticHandler(sc []StaticContent, path, root string) []StaticContent {
	return append(sc, StaticContent{
		Path: path,
		Root: root,
	})
}
