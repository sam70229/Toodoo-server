package routes

import (
    "encoding/json"
    "net/http"

    "github.com/go-chi/chi"

	"Toodoo/model"
)


func (s *Server) GetTodos() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		start_time := r.URL.Query().Get("start_time")
		end_time := r.URL.Query().Get("end_time")

		s.logger.Infow("GetTodos", "query", r.URL.Query())

		todos, err := s.apistore.GetTodos(ctx, start_time, end_time)
		if err != nil {
            s.logger.Fatalw("GetTodos", "error", err)
		}

		response := ApiResponse{todos}
		d, _ := json.Marshal(response)
		w.Write(d)
	}
}

func (s *Server) GetTodoById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		uid := chi.URLParam(r, "uid")

		todo, err := s.apistore.GetTodoById(ctx, uid)

		if err != nil {
            s.logger.Fatalw("GetTodoById", "error", err)
		}

		response := ApiResponse{todo}
		d, _ := json.Marshal(response)
		w.Write(d)
	}
}

func (s *Server) AddTodo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var request map[string][]model.Todo
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&request)
		if err != nil {
            s.logger.Fatalw("AddTodo", "error", err)
		}
		s.logger.Infow("AddTodo", "request", request)
		// checkError(err)

        uid, err := s.apistore.AddTodo(ctx, request["todos"][0])

		response := ApiResponse{uid}
		d, _ := json.Marshal(response)
		w.Write(d)
	}
}
