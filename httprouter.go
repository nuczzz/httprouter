package httprouter

import (
	"net/http"
)

// Handler handler of http.
type Handler func(resp http.ResponseWriter, req *http.Request)

// Router router interface definition.These methods like Get, Head
// are functions to add route handler with path and handler,and cover
// it if old handler already exist.
type Router interface {
	// Get add get handler
	Get(path string, handler Handler)
	// Head add head handler
	Head(path string, handler Handler)
	// Post add post handler
	Post(path string, handler Handler)
	// Put add put handler
	Put(path string, handler Handler)
	// Patch add patch handler
	Patch(path string, handler Handler)
	// Delete add delete handler
	Delete(path string, handler Handler)
	// Connect add connect handler
	Connect(path string, handler Handler)
	// Options add options handler
	Options(path string, handler Handler)
	// Trace add trace handler
	Trace(path string, handler Handler)

	// Match return the handler by method and path.if not match
	// any handler, NotFound will be return.
	Match(method, path string) Handler
}

// NewRouter return a new router instance.
func NewRouter() Router {
	r := &router{
		routers: make(map[string]map[string]Handler),
	}
	for _, method := range RouterMethods {
		r.routers[method] = make(map[string]Handler)
	}
	return r
}
