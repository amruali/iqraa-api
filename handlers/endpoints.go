package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func (s *Server) setupEndpoints(router *chi.Mux) {
	router.Route("/api", func(router chi.Router) {
		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode("Ok")
		})
		router.Get("/users", s.SampleUser)
	})
}
