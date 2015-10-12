package router

import "net/http"

type Route struct {
	Method, Path string
	Handler http.HandlerFunc
}

type Router struct {
	routes map[string][]*Route
}

func New() *Router {
	return &Router{routes:make(map[string][]*Route)}
}

func (router *Router) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	for _, route := range(router.routes[r.Method]) {
		if (route.Path == r.URL.Path) {
			route.Handler(rw, r)
			return
		}
	}
}

func (router *Router) Register(method, path string, handler http.HandlerFunc)  {
	router.routes[method] = append(router.routes[method], &Route{Method:method, Path:path, Handler:handler})
}