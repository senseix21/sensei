package router

import (
	"net"
	"regexp"
)

// Hnadler func defines the RequestHandler
type handlerFunc func(net.Conn, map[string]string)

// Route struct to store the route patterns
type Route struct {
	Method  string
	Pattern string
	Regex   *regexp.Regexp
	Params  []string
	Handler handlerFunc
}

// Router struct to store routes
type Router struct {
	routes []*Route
}

// NewRouter creates a new router
func NewRouter() *Router {
	return &Router{}
}

// Hnadle registers a route with the handler function
func (r *Router) Handle(method, pattern string, handler handlerFunc) {
	paramNames := []string{}
	regexPattern := regexp.MustCompile(`:[a-zA-Z0-9_]+`)

	//Extract parameter names
	matches := regexPattern.FindAllString(pattern, -1)
	for _, match := range matches {
		paramNames = append(paramNames, match[1:])
	}

	// Replace parameterized route with regex
	regexStr := regexPattern.ReplaceAllString(pattern, `([^/]+)`)
	regex := regexp.MustCompile("^" + regexStr + "$")

	// Add route
	route := &Route{Method: method, Pattern: pattern, Regex: regex, Params: paramNames, Handler: handler}
	r.routes = append(r.routes, route)
}

// Serve process incming requests
func (r *Router) Serve(conn net.Conn, method, path string) {
	for _, route := range r.routes {
		if route.Method == method && route.Regex.MatchString(path) {
			matches := route.Regex.FindStringSubmatch(path)
			params := map[string]string{}

			//Map path with params
			for i, match := range route.Params {
				params[match] = matches[i+1]
			}
			route.Handler(conn, params)
			return
		}
	}
	conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
}
