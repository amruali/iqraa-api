package handlers

import (
	"iqraa-api/domain"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

// Get Profile
func (s *Server) GetProfile() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := chi.URLParam(r, "username")
		if len(username) < 1 {
			resposneWithJson(w, map[string]string{"error": "username is required"}, http.StatusBadRequest)
		}

		if user, err := s.domain.GetCashedStrings(username, domain.Profile{}); err == nil {
			resposneWithJson(w, user, http.StatusOK)
			return
		} else {
			log.Printf("data hasn't cashed %v", err)
			user, err := s.domain.GetUserInfo(username)
			if err != nil {
				resposneWithJson(w, map[string]string{"error": err.Error()}, http.StatusNotFound)
				return
			}
			resposneWithJson(w, user, http.StatusOK)
			go s.domain.SetCashedStrings(username, ToString(user))
		}
	})
}

// Upvote   (setting file Mongo || setting Column || upvotes column)

// Downvote (setting file Mongo || setting Column || upvotes downvotes)

// Recommended (setting file Mongo || setting Column || recommended downvotes)

// Request to Upload Book (token + page to Upload Book as a user) ==> email verification needed

// Report book

// Report author

// Report publisher

// Report User  (USER'S GROUP OF REPORRTED USERS ==> COLUMN )

// Propose Feature / Complain problem  ==> Table

// Add Book to shelves  (USER GROUP OF BOOKS == COLUMN)
