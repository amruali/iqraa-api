package handlers

import (
	"context"
	"encoding/json"
	"iqraa-api/domain"
	"net/http"
)

type (
	authenticationResponse struct {
		User  *domain.User     `json:"user"`
		Token *domain.JwtToken `json:"token"`
	}
	PayloadValidation interface {
		IsValid() (bool, map[string]string)
	}
)

func ValidatePayload(next http.HandlerFunc, payload PayloadValidation) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			BadRequest(w, err)
			return
		}

		defer r.Body.Close()

		if valid, err := payload.IsValid(); !valid {
			resposneWithJson(w, err, http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), "payload", payload)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func resposneWithJson(w http.ResponseWriter, payload interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if payload == nil {
		payload = map[string]string{}
	}
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func BadRequest(w http.ResponseWriter, err error) {
	response := map[string]string{"error": err.Error()}
	resposneWithJson(w, response, http.StatusBadRequest)
}

func unAuthorizedResponse(w http.ResponseWriter) {
	response := map[string]string{"error": "UnAuthorized"}
	resposneWithJson(w, response, http.StatusUnauthorized)
}

func forbiddenResponse(w http.ResponseWriter) {
	response := map[string]string{"error": "Forbidden"}
	resposneWithJson(w, response, http.StatusForbidden)
}

func ToString(value interface{}) string {
	bytes, _ := json.Marshal(value)
	return string(bytes)
}
