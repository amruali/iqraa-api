package models

import "time"

type Review struct {
	ReviewID     int32     `json:"review_id"`
	BookID       int32     `json:"book_id"`
	Description  string    `json:"review_desc"`
	CreateUserID uint32    `json:"create_user_id"`
	UpdateUserID uint32    `json:"update_user_id"`
	StarsCount   uint8     `json:"stars_count"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
