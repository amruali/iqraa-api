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
		router.Get("/", s.SampleAllUserBooks) // test

		// user-group
		router.Route("/users", func(r chi.Router) {
			r.Post("/login", s.Login())
			r.Post("/register", s.Register())
		})

		// book-group
		router.Route("/books", func(r chi.Router) {
			r.Post("/", s.CreateBook())

			r.Route("/{book_id}", func(r chi.Router) {
				//r.Use(s.GetBookFromCTX)
				//r.Use(s.withOwner("todo"))
				r.Put("/", s.UpdateBook())
				r.Get("/", s.GetBook)
				r.Delete("/", s.DeleteBook())
				//r.Patch("/", s.UpdateTodo())

			})

		})

	})
}
