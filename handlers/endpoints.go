package handlers

import (
	"github.com/go-chi/chi"
)

func (s *Server) setupEndpoints(router *chi.Mux) {
	router.Route("/api", func(router chi.Router) {
		/*
			router.Get("/", func(w http.ResponseWriter, r *http.Request) {
				json.NewEncoder(w).Encode("Ok")
			})
		*/
		router.Route("/users", func(r chi.Router) {
			//router.Get("/", s.SampleUser)
			r.Post("/login", s.Login())
			r.Post("/register", s.Register())
		})

	})
}
