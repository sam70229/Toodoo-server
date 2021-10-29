package routes

import (


	"github.com/go-chi/chi"
	"go.uber.org/zap"

	"Toodoo/api"
)

var apiVersion = "/api/v1"

// func init() {
// 	register("POST", apiVersion + "/todo/add", controller.AddTodo, nil)

// 	register("GET", apiVersion + "/todos", controller.GetTodos, nil)
// 	register("GET", apiVersion + "/todo/{uid}", controller.GetTodoById, nil)
// }

type Server struct {
	logger *zap.SugaredLogger
	router chi.Router
	apistore api.APIStore
}

type ApiResponse struct {
	Data interface{} `json:"data"`
}

func NewRouter(router chi.Router, apistore api.APIStore) error {

	s := &Server{
		logger: zap.S().With("package", "routes"),
		router: router,
		apistore: apistore,
	}

	s.router.Route(apiVersion, func(r chi.Router) {
		r.Get("/todos", s.GetTodos())
		r.Get("/todo/{uid}", s.GetTodoById())
		r.Post("/todo/add", s.AddTodo())
		r.Get("/categories", s.GetCategories())
		r.Post("/category/add", s.AddCategory())
		r.Patch("/", s.UpdateTodo())
	})

	return nil
}

