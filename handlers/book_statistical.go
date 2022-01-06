package handlers

import (
	"net/http"
)

func (s *Server) TopDownloads() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		books, err := s.domain.GetTopDownloadedBooks()
		if err != nil {
			resposneWithJson(w, map[string]string{"error": "An error happened"}, http.StatusInternalServerError)
			return
		}
		go s.domain.SetCashedTopDownloads(ToString(books))
		resposneWithJson(w, books, http.StatusOK)
	})
}
