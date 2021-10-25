package routes

import (
	"encoding/json"
	"net/http"

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
	Data interface{}
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
	})

	return nil
}

func (s *Server) GetCategories() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		s.logger.Infow("GetCategories", "query", r.URL.Query())
		categories, err := s.apistore.GetCategories(ctx)
		if err != nil {

		}

		response := ApiResponse{categories}
		d, _ := json.Marshal(response)
		w.Write(d)
	}
}