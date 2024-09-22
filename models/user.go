package models

import "time"

type User struct {
	Id             int64                  `json:"user_id"`
	Username       string                 `json:"username"`
	Email          string                 `json:"email"`
	EmailConfirmed string                 `json:"email_confirmed"`
	HashedPassword []byte                 `json:"-"`
	UserTypeID     int8                   `json:"user_type_id"`
	CreatedAt      time.Time              `json:"created_at"`
	UpdatedAt      time.Time              `json:"updated_at"`
	FirstName      string                 `json:"first_name"`
	LastName       string                 `json:"last_name"`
	Image          string                 `json:"image"`
	Settings       map[string]interface{} `json:"likes"`
}
