package router

import "net/http"

type Route struct {
	Method, Path       string
	MultiPartVariables bool
	Handler            http.HandlerFunc
}

type Router struct {
	routes map[string][]*Route
}

func New() *Router {
	return &Router{routes: make(map[string][]*Route)}
}

func (router *Router) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	for _, route := range router.routes[r.Method] {
		if route.Path == r.URL.Path {
			route.Handler(rw, r)
			return
		}
	}
}

func (router *Router) RegisterNew(method, path string, handler http.HandlerFunc) {
	router.RegisterRoute(&Route{Method: method, Path: path, Handler: handler})
}

func (router *Router) RegisterRoute(route *Route) {
	router.routes[route.Method] = append(router.routes[route.Method], route)
}

func (router *Router) RegisterBatch(routes []*Route) {
	for _, route := range routes {
		router.RegisterRoute(route)
	}
}
