package domain

import "time"

type Quote struct {
	ID           uint8     `json:"quote_id"`
	BookID       uint32    `json:"book_id"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	CreateUserID uint32    `json:"create_user_id"`
	UpdateUserID uint32    `json:"update_user_id"`
}
