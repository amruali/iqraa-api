package handlers

import (
	"github.com/go-chi/chi"
	_"encoding/json"
	"net/http"
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
	router.Route("/api/v1", func(router chi.Router){
		router.Route("/users", func(r chi.Router) {
			router.Get("/", s.SampleUser)
			r.Post("/login", s.Login())
			r.Post("/register", s.Register())
		})
	
	})
}
