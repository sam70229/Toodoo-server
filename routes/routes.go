package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"Toodoo/controller"
)

type Route struct {
	Method string
	Pattern string
	Handler http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

var routes []Route

var apiVersion = "/api/v1"

func init() {
	register("POST", apiVersion + "/todo/add", controller.AddTodo, nil)

	register("GET", apiVersion + "/todos", controller.GetTodos, nil)
	register("GET", apiVersion + "/todo/{uid}", controller.GetTodoById, nil)
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	for _, route := range routes {
		r.Methods(route.Method).
		Path(route.Pattern).
		Handler(route.Handler)
		if route.Middleware != nil {
			r.Use(route.Middleware)
		}
	}
	return r
}

func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, Route{method, pattern, handler, middleware})
}