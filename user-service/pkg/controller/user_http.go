package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maestre3d/quark-demo/user-service/internal/application"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type UserHTTP struct {
	App *application.User
}

func (c *UserHTTP) SetEndpoints(r *mux.Router) {
	r.Path("/user").Methods(http.MethodPost).HandlerFunc(c.create)
}

func (c *UserHTTP) create(w http.ResponseWriter, r *http.Request) {
	id, err := gonanoid.New()
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		})
		return
	}
	err = c.App.Create(r.Context(), id, r.PostFormValue("username"), r.PostFormValue("email"))
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(struct {
			Message string `json:"message"`
		}{
			Message: err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		ID string `json:"user_id"`
	}{
		ID: id,
	})
}
