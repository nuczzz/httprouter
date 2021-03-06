package httprouter

import (
	"net/http"
)

const (
	MethodGet     = http.MethodGet
	MethodHead    = http.MethodHead
	MethodPost    = http.MethodPost
	MethodPut     = http.MethodPut
	MethodPatch   = http.MethodPatch
	MethodDelete  = http.MethodDelete
	MethodConnect = http.MethodConnect
	MethodOptions = http.MethodOptions
	MethodTrace   = http.MethodTrace
)

var RouterMethods = []string{
	MethodGet,
	MethodHead,
	MethodPost,
	MethodPut,
	MethodPatch,
	MethodDelete,
	MethodConnect,
	MethodOptions,
	MethodTrace,
}

type router struct {
	// routers map of handler with method and path.
	// key1-method, key2-path
	routers map[string]map[string]Handler

	// ifNotMatch handler of not match, default is http.NotFound
	ifNotMatch Handler
}

func newRouter() *router {
	r := &router{
		routers: make(map[string]map[string]Handler),
	}
	for _, method := range RouterMethods {
		r.routers[method] = make(map[string]Handler)
	}
	return r
}

func (r *router) SetIfNotMatch(handler Handler) {
	r.ifNotMatch = handler
}

func (r *router) Get(path string, handler Handler) {
	r.routers[MethodGet][path] = handler
}

func (r *router) Head(path string, handler Handler) {
	r.routers[MethodHead][path] = handler
}

func (r *router) Post(path string, handler Handler) {
	r.routers[MethodPost][path] = handler
}

func (r *router) Put(path string, handler Handler) {
	r.routers[MethodPut][path] = handler
}

func (r *router) Patch(path string, handler Handler) {
	r.routers[MethodPatch][path] = handler
}

func (r *router) Delete(path string, handler Handler) {
	r.routers[MethodDelete][path] = handler
}

func (r *router) Connect(path string, handler Handler) {
	r.routers[MethodConnect][path] = handler
}

func (r *router) Options(path string, handler Handler) {
	r.routers[MethodOptions][path] = handler
}

func (r *router) Trace(path string, handler Handler) {
	r.routers[MethodTrace][path] = handler
}

// Match return the handler by method and path.if not match
// any handler, NotFound will be return.
func (r *router) match(method, path string) Handler {
	if v1, ok := r.routers[method]; ok {
		if handle, ok := v1[path]; ok {
			return handle
		}
	}
	if r.ifNotMatch != nil {
		return r.ifNotMatch
	}
	return http.NotFound
}

// ServeHTTP implement http.Handler
func (r *router) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	r.match(req.Method, req.URL.Path)(resp, req)
}
