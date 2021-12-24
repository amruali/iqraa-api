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
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		resposneWithJson(w, struct {
			Ok string
		}{Ok: "Welcome, bro"}, http.StatusOK)
	})
	router.Route("/api/v1", func(router chi.Router) {
		router.Get("/", s.SampleAllUserBooks) // test

		// user-group
		router.Route("/users", func(r chi.Router) {
			r.Post("/login", s.Login())
			r.Post("/register", s.Register())

			r.Route("/{username}/reviews", func(r chi.Router) {
				//r.Get("/", s.UserReviews())
			})
		})

		// books-group
		router.Route("/books", func(r chi.Router) {
			r.Get("/", s.GetBook)
			//r.Get("/most_read", s.GetBook())
			//r.Get("/new", s.GetBook())
			//r.Get("/yours", s.GetBook())
			//r.Get("/{book_id}/reviews", s.BookReviews())
			// Admin-Only
			r.Post("/", s.CreateBook())
			r.Route("/{book_id}", func(r chi.Router) {
				r.Use(s.IsAdmin)
				r.Put("/", s.UpdateBook())
				r.Delete("/", s.DeleteBook())
				//r.Patch("/", s.UploadBookImage())
			})
		})

		router.Route("/reviews", func(r chi.Router) {
			r.Use(s.IsAuthenticated)
			//r.Post("/", s.ReviewBook())
			r.Route("/{review_id}", func(r chi.Router) {
				//r.Put("/", s.UpdateReview())
				//r.Delete("/", s.DeleteReview())
			})
		})

		router.Route("/authors", func(r chi.Router) {
			//r.Get("/", s.GetAuthors())
			r.Route("/{author_id}", func(r chi.Router) {
				//r.Get("/", s.AuthorInfo())
				r.Route("/", func(r chi.Router) {
					r.Use(s.IsAdmin)
					//r.Put("", s.UpdateAuthorInfo())
				})
			})
		})
	})
}
