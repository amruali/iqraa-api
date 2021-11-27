package domain

import "time"

type Review struct {
	ReviewID   int32     `json:"review_id"`
	BookID     int32     `json:"book_id"`
	ReviewDesc string    `json:"review_desc"`
	UserID     int32     `json:"string"`
	StarsCount uint8     `json:"stars_count"`
	CreateAt   time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}