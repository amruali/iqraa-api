package domain

import "time"

type Book struct {
	BookID           int64     `json:"book_id"`
	BookName         string    `json:"book_name"`
	BookISBN         string    `json:"isbn"`
	BookPublishDate  time.Time `json:"publish_date"`
	BookPublisherID  int32     `json:"publisher_id"`
	BookTypeID       int16     `json:"book_type_id"`
	BookTypeDetailID int32     `json:"book_type_detail_id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	CreateUserID     int64     `json:"create_user_id"`
	UpdateUserID     int64     `json:"update_user_id"`
}


