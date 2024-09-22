package dtos

import (
	"time"
)

type Profile struct {
	Username   string                 `json:"username"`
	Email      string                 `json:"email"`
	UserTypeID int8                   `json:"user_type_id"`
	CreatedAt  time.Time              `json:"created_at"`
	UpdatedAt  time.Time              `json:"updated_at"`
	FirstName  string                 `json:"first_name"`
	LastName   string                 `json:"last_name"`
	Image      string                 `json:"image"`
	Settings   map[string]interface{} `json:"likes"`
}