package handlers

import (
	"fmt"
	"iqraa-api/domain"
	"net/http"
)

func (s *Server) CreateBook() http.HandlerFunc {
	var payload domain.CreateBookPayload
	return ValidatePayload(func(w http.ResponseWriter, r *http.Request) {
		book, err := s.domain.CreateBook(payload)
		if err != nil {
			fmt.Println(err)
			BadRequest(w, err)
			return
		}
		resposneWithJson(w, book, http.StatusCreated)
	}, &payload)
}
