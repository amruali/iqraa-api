package handlers

import (
	_ "encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func favHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/favicon.ico" {
			w.WriteHeader(http.StatusNoContent)
			next.ServeHTTP(w, r)
		}
		next.ServeHTTP(w, r)
	})
}

func (s *Server) setupEndpoints(router *chi.Mux) {
	router.Route("/api/v1", func(router chi.Router) {
		router.Get("/", s.SampleUser) // test

		// user-group
		router.Route("/users", func(r chi.Router) {
			r.Post("/login", s.Login())
			r.Post("/register", s.Register())
		})

		// book-group
		router.Route("/books", func(r chi.Router) {
			r.Post("/new", s.CreateBook())
		})

	})
}
