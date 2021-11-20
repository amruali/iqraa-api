package domain

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Id             int64     `json:"user_id"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	HashedPassword []byte    `json:"-"`
	UserTypeID     int8      `json:"user_type_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Image          string    `json:"image"`
}

type JwtToken struct {
	AccessToken string    `json:"token"`
	ExpiresAt   time.Time `json:"expires_at`
}

func (u *User) GenerateJwtToken() (*JwtToken, error) {
	jwtToken := jwt.New(jwt.GetSigningMethod("HS256"))

	expiresAt := time.Now().Add(time.Hour * 24 * 7) // a Single Week

	jwtToken.Claims = jwt.MapClaims{
		"id":       u.Id,
		"username": u.Username,
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

func (d *Domain) GetUserByUserID(id int64) (*User, error) {
	user, err := d.DB.UserRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
