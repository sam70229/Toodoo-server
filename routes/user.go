package routes

import (
	"Toodoo/model"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type UserRequest struct {
	DeviceId string `json:"deviceId"`
	UserId string `json:"userId"`
	Os string `json:"os"`
	Os_ver string `json:"os_ver"`
}

func (s *Server) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		s.logger.Info("Get login called")
		token := tokenGenerator()

		var userRequest UserRequest
		json.NewDecoder(r.Body).Decode(&userRequest)

		user := &model.User{
			UserId: userRequest.UserId,
			DeviceId: userRequest.DeviceId,
			Token: token,
			Create_at: time.Now().Unix(),
			Os: userRequest.Os,
			Os_ver: userRequest.Os_ver,
		}		
		
		s.logger.Infow("Login", "request", userRequest)
		s.logger.Infow("Login", "user data", user)
		s.apistore.RegisterUser(ctx, user)

		w.Header().Add("Cookie", "token=" + token)
		RenderJSONResponse(w, http.StatusOK, nil)
	}
}

func tokenGenerator() string {
    b := make([]byte, 32)
    rand.Read(b)
    return fmt.Sprintf("%x", b)
}
