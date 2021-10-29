package routes

import (
	"Toodoo/model"
	"encoding/json"
	"net/http"
)

func (s *Server) GetCategories() http.HandlerFunc {
   return func(w http.ResponseWriter, r *http.Request) {
       ctx := r.Context()

       s.logger.Infow("GetCategories", "query", r.URL.Query())

       categories, err := s.apistore.GetCategories(ctx)
       if err != nil {
           s.logger.Fatalw("GetCategories", "error", err)
       }
       response := ApiResponse{categories}
       d, _ := json.Marshal(response)
       s.logger.Infow("GetCategories", "data", d)
       w.Write(d)
   }
}

func (s *Server) AddCategory() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        ctx := r.Context()
        var request map[string][]model.Category
        decoder := json.NewDecoder(r.Body)
        err := decoder.Decode(&request)
        if err != nil {
            s.logger.Fatalw("AddCategory", "error", err)
        }

        uid, err := s.apistore.AddCategory(ctx, request["categories"][0])

        RenderJSONResponse(w, http.StatusOK, uid)
    }
}
