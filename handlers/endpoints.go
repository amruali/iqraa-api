package handlers

import (
	"github.com/go-chi/chi"
	"encoding/json"
	"net/http"
)


func favHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.path == "/favicon.ico"{
			w.WriteHeader(http.StatusNoContent)
			next.ServeHTTP(w, r)
		}
		next.ServeHTTP(w, r)
	})
}

func (s *Server) setupEndpoints(router *chi.Mux) {	
	router.Route("", func(router chi.Router){
		router.Use(favHandler);
		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode("Ok")
		})
		router.Route("/api", func(router chi.Router) {
			router.Route("/users", func(r chi.Router) {
				//router.Get("/", s.SampleUser)
				r.Post("/login", s.Login())
				r.Post("/register", s.Register())
			})
	
		})
	})
}
