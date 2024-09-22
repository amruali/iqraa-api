package handlers

import (
	"context"
	"iqraa-api/models"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

type JwtToken struct {
	AccessToken string    `json:"token"`
	ExpiresAt   time.Time `json:"expires_at`
}

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

func GenerateJwtToken(u *models.User) (*JwtToken, error) {

	jwtToken := jwt.New(jwt.GetSigningMethod("HS256"))

	expiresAt := time.Now().Add(time.Hour * 24 * 7) // a Single Week

	jwtToken.Claims = jwt.MapClaims{
		"id":       u.Id,
		"username": u.Username,
		"role":     u.UserTypeID,
		"exp":      expiresAt.Unix(),
	}

	accessToken, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	return &JwtToken{
		AccessToken: accessToken,
		ExpiresAt:   expiresAt,
	}, nil
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

/*
	func (s *Server) IsCashed(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			books, err := s.domain.GetCashedTopDownloads()
			if err == nil {
				resposneWithJson(w, books, http.StatusOK)
				return
			} else {
				log.Printf("redis failed to return data because %v", err)
				next.ServeHTTP(w, r)
			}
		})
	}
*/
func (s *Server) IsCashed(Key string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch Key {
			case "top_downloads":
				if books, err := s.domain.GetCashedStrings(Key, []models.Book{}); err == nil {
					resposneWithJson(w, books, http.StatusOK)
					return
				} else {
					log.Printf("data has not cashed yet")
					ctx := context.WithValue(r.Context(), Key, Key)
					next.ServeHTTP(w, r.WithContext(ctx))
				}
			default:
				log.Println("enta developer t3ban")
			}
		})
	}
}
