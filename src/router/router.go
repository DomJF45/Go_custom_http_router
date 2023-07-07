package router

import (
	"net/http"

	routererrors "custom_http_router/src/router_errors"
)

type Router struct {
	tree *tree
}

type route struct {
	methods []string
	path    string
	handler http.Handler
}

var tmpRoute = &route{}

func NewRouter() *Router {
	return &Router{
		tree: NewTree(),
	}
}

func (r *Router) Methods(methods ...string) *Router {
	tmpRoute.methods = append(tmpRoute.methods, methods...)
	return r
}

func (r *Router) Handler(path string, handler http.Handler) {
	tmpRoute.handler = handler
	tmpRoute.path = path
	r.Handle()
}

func (r *Router) Handle() {
	r.tree.Insert(tmpRoute.methods, tmpRoute.path, tmpRoute.handler)
	tmpRoute = &route{}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	path := req.URL.Path
	result, err := r.tree.Search(method, path)
	if err != nil {
		status := handleErr(err)
		w.WriteHeader(status)
		return
	}
	h := result.actions.handler
	h.ServeHTTP(w, req)
}

func handleErr(err error) int {
	var status int
	switch err {
	case routererrors.ErrMethodNotAllowed:
		status = http.StatusMethodNotAllowed
	case routererrors.ErrNotFound:
		status = http.StatusNotFound
	}
	return status
}
