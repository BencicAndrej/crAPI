package router

import (
	"net/http"
	"log"
	"regexp"
)

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
//		if !route.MultiPartVariables {
//			reg, err := regexp.Compile("{.*}")
//
//			regexdString := reg.ReplaceAllString(route.Path, "")
//
//			parts := strings.Split(route.Path, "/")
//		} else {
//
//		}
		matched, err := regexp.Match(route.Path, []byte(r.URL.Path))
		if err != nil {
			log.Print(err)
		}

		if matched {
			route.Handler(rw, r)
			return
		}
	}
}

func (router *Router) RegisterNew(method, path string, handler http.HandlerFunc) {
	router.RegisterRoute(&Route{Method: method, Path: path, Handler: handler})
}

func (router *Router) RegisterBatch(routes []*Route) {
	for _, route := range routes {
		router.RegisterRoute(route)
	}
}

func (router *Router) RegisterRoute(route *Route) {
	router.routes[route.Method] = append(router.routes[route.Method], route)
	log.Printf("Registered route: %v %v.\n", route.Method, route.Path)
}


func (router *Router) GetRouteMap() map[string][]*Route {
	return router.routes
}