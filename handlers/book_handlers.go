package handlers

import (
	"iqraa-api/domain"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// Add Upload pdf file // Cover
func (s *Server) CreateBook() http.HandlerFunc {
	var payload domain.CreateBookPayload
	return ValidatePayload(func(w http.ResponseWriter, r *http.Request) {
		book, err := s.domain.CreateBook(payload)
		if err != nil {
			BadRequest(w, err)
			return
		}
		resposneWithJson(w, book, http.StatusCreated)
	}, &payload)
}

// Update Book - Edit book info
func (s *Server) UpdateBook() http.HandlerFunc {
	var payload domain.UpdateBookPayload
	return ValidatePayload(func(w http.ResponseWriter, r *http.Request) {
		bookID := chi.URLParam(r, "book_id")
		id, err := strconv.ParseInt(bookID, 10, 64)

		if err != nil {
			resposneWithJson(w, map[string]string{"error": "book is not found"}, http.StatusNotFound)
			return
		}
		err = s.domain.UpdateBook(payload, id)
		if err != nil {
			BadRequest(w, err)
			return
		}
		resposneWithJson(w, struct {
			Ok string
		}{Ok: "Successfully Updated"}, http.StatusCreated)
	}, &payload)
}

func (s *Server) GetBook(w http.ResponseWriter, r *http.Request) {
	bookID := chi.URLParam(r, "book_id")
	if len(bookID) < 1 {
		resposneWithJson(w, map[string]string{"error": "book id is required"}, http.StatusBadRequest)
	}
	id, err := strconv.ParseInt(bookID, 10, 64)

	if err != nil {
		resposneWithJson(w, map[string]string{"error": "book is not found"}, http.StatusNotFound)
		return
	}

	book, err := s.domain.GetBook(id)
	if err != nil {
		resposneWithJson(w, map[string]string{"error": err.Error()}, http.StatusNotFound)
		return
	}

	resposneWithJson(w, book, http.StatusOK)

}

func (s *Server) DeleteBook() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bookID := chi.URLParam(r, "book_id")
		if len(bookID) < 1 {
			resposneWithJson(w, map[string]string{"error": "book id is required"}, http.StatusBadRequest)
		}
		id, err := strconv.ParseInt(bookID, 10, 64)

		if err != nil {
			resposneWithJson(w, map[string]string{"error": "book is not found"}, http.StatusNotFound)
			return
		}

		err = s.domain.DeleteBook(id)
		if err != nil {
			resposneWithJson(w, map[string]string{"error": err.Error()}, http.StatusNotFound)
			return
		}

		resposneWithJson(w, map[string]string{"Ok": "Successfully Deleted"}, http.StatusOK)

	})
}

// Update Book URL

// Upload New PDF Copy

// Delete Book
