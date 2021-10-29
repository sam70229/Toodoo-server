package routes

import (
	"bytes"
	"encoding/json"
	"net/http"

)

func RenderJSONResponse(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	var apiResponse ApiResponse
	apiResponse.Data = v
	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(apiResponse); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(code)
	}
	_, _ = w.Write(b.Bytes())
}