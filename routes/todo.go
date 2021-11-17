package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"Toodoo/model"
)


func (s *Server) GetTodos() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		fmt.Println(r.Header)
		fmt.Println(r.Cookie("user"))

		start_time := r.URL.Query().Get("start_time")
		end_time := r.URL.Query().Get("end_time")
		if r.URL.Query().Get("category") != "" {
			category := r.URL.Query().Get("category")
			s.logger.Infow("GetTodos", "category", category)

			todos, err := s.apistore.GetTodosByCategory(ctx, category)
			if err != nil {
				s.logger.Fatalw("GetTodos", "error", err)
			}
			RenderJSONResponse(w, http.StatusOK, todos)
		} else {
			s.logger.Infow("GetTodos", "query", r.URL.Query())

			todos, err := s.apistore.GetTodos(ctx, start_time, end_time)
			if err != nil {
				s.logger.Fatalw("GetTodos", "error", err)
			}
	
			RenderJSONResponse(w, http.StatusOK, todos)
		}
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

		var request []model.Todo
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&request)
		if err != nil {
            s.logger.Fatalw("AddTodo", "error", err)
		}
		s.logger.Infow("AddTodo", "request", request)
		// checkError(err)

        uid, err := s.apistore.AddTodo(ctx, request[0])

		RenderJSONResponse(w, http.StatusOK, uid)

	}
}

func (s *Server) UpdateTodo() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		// ctx := r.Context()
		
	}

}