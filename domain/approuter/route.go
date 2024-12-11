package approuter

import "net/http"

type Route struct {
	path    string
	handler http.HandlerFunc
}

func (r Route) Path() string {
	return r.path
}

func (r Route) HandlerFunc() http.HandlerFunc {
	return r.handler
}

func NewRoute(path string, handler http.HandlerFunc) Route {
	return Route{path, handler}
}
