package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) SampleUser(w http.ResponseWriter, r *http.Request) {
	user, err := s.domain.DB.UserRepo.GetByID(1)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(user.Email)
}
