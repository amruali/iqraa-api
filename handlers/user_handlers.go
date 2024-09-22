package handlers

import (
	"encoding/json"
	"fmt"
	"iqraa-api/dtos"
	"net/http"
)

func (s *Server) SampleAllUserBooks(w http.ResponseWriter, r *http.Request) {

	books, err := s.domain.DB.ReviewRepo.GetByUserID(1)

	if err != nil {
		fmt.Println(err, 1)
	}

	fmt.Println(books, 2)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(books)
}

func (s *Server) Register() http.HandlerFunc {
	var payload dtos.RegisterPayload
	return ValidatePayload(func(w http.ResponseWriter, r *http.Request) {

		user, err := s.domain.Register(payload)
		if err != nil {
			fmt.Println(err)
			BadRequest(w, err)
			return
		}

		// Generate Token
		token, err := GenerateJwtToken(user)
		if err != nil {
			BadRequest(w, err)
			return
		}

		resposneWithJson(w, &authenticationResponse{
			User:  user,
			Token: token,
		}, http.StatusCreated)
	}, &payload)
}

func (s *Server) Login() http.HandlerFunc {
	var payload dtos.LoginPayload
	return ValidatePayload(func(w http.ResponseWriter, r *http.Request) {

		user, err := s.domain.Login(payload)

		if err != nil {
			BadRequest(w, err)
			return
		}

		// Generate Token
		token, err := GenerateJwtToken(user)
		if err != nil {
			BadRequest(w, err)
			return
		}

		resposneWithJson(w, &authenticationResponse{
			User:  user,
			Token: token,
		}, http.StatusOK)

	}, &payload)

}
