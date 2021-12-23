package handlers

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

func stripBearerPrefixToken(token string) (string, error) {
	bearer := "BEARER"
	if len(token) > len(bearer) && strings.ToUpper(token[0:len(bearer)]) == bearer {
		return token[len(bearer)+1:], nil
	}
	return token, nil
}

var authHeaderExtractor = &request.PostExtractionFilter{
	//Extractor : request.HeaderExtractor{"Authorization", "key1", key2}
	Extractor: request.HeaderExtractor{"Authorization"},
	Filter:    stripBearerPrefixToken,
}

var authExtractor = &request.MultiExtractor{
	authHeaderExtractor,
}

func ParseToken(r *http.Request) (*jwt.Token, error) {
	token, err := request.ParseFromRequest(r, authExtractor, func(t *jwt.Token) (interface{}, error) {
		b := []byte(os.Getenv("JWT_SECRET"))
		return b, nil
	})

	return token, err
}

func (s *Server) IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse Token
		token, err := ParseToken(r)
		if err != nil {
			unAuthorizedResponse(w)
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := int64(claims["id"].(float64))
			ctx := context.WithValue(r.Context(), "userID", userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			unAuthorizedResponse(w)
			return
		}
	})
}

func (s *Server) IsAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse Token
		token, err := ParseToken(r)
		if err != nil {
			unAuthorizedResponse(w)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := int64(claims["id"].(float64))
			userRole := int64(claims["role"].(float64))

			if userRole != 1 {
				unAuthorizedResponse(w)
				return
			}
			ctx := context.WithValue(r.Context(), "userID", userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			unAuthorizedResponse(w)
			return
		}
	})
}
