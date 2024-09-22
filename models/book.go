package models

import "time"

type Book struct {
	ID               int64     `json:"id"`
	BookName         string    `json:"book_name"`
	ISBN             string    `json:"isbn"`
	BookAuthorID     int32     `json:"author_id"`
	PublishYear      int32     `json:"publish_year"`
	PublisherID      int32     `json:"publisher_id"`
	BookTypeID       int32     `json:"book_type_id"`
	BookTypeDetailID int32     `json:"book_type_detail_id"`
	DownloadsNumber  int32     `json:"downloads_number"`
	Brief            string    `json:"brief"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	CreateUserID     int32     `json:"create_user_id"`
	UpdateUserID     int32     `json:"update_user_id"`
}