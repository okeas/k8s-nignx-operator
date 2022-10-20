package config

import (
	"github.com/gorilla/mux"
	"net/http"
)

var MyRouter *mux.Router


type RouteBuilder struct {
	route *mux.Route
}

func init() {
	MyRouter = mux.NewRouter()
}

func NewRouteBuilder() *RouteBuilder {
	return &RouteBuilder{
		route: MyRouter.NewRoute(),
	}
}

func (r *RouteBuilder) SetPath(path string, exact bool) *RouteBuilder {
	if exact {
		r.route.Path(path)
	} else {
		r.route.PathPrefix(path)
	}
	return r
}

func (r *RouteBuilder) SetHost(host string, set bool) *RouteBuilder {
	if set {
		r.route.Host(host)
	}
	return r
}

func (r *RouteBuilder) Build(handler http.Handler) {
	r.route.
		Methods("GET", "POST", "PUT", "DELETE", "OPTIONS").
		Handler(handler)
}


