package router

import (
	"fmt"
	"net"
)

// Hnadler func defines the RequestHandler
type handlerFunc func(net.Conn, string)

// Router struct to store routes
type Router struct {
	routes map[string]handlerFunc
}

// NewRouter creates a new router
func NewRouter() *Router {
	return &Router{routes: make(map[string]handlerFunc)}
}

// Hnadle registers a route with the handler function
func (r *Router) Handle(method, path string, handler handlerFunc) {
	key := fmt.Sprintf("%s %s", method, path)
	r.routes[key] = handler
}

// Serve process incming requests
func (r *Router) Serve(conn net.Conn, method, path string) {
	key := fmt.Sprintf("%s %s", method, path)
	handler, ok := r.routes[key]
	if !ok {
		fmt.Fprintf(conn, "HTTP 1.1 404 Not Found\r\n\r\n")
		return
	}
	handler(conn, path)
}
