package domain

import (
	"time"
)

type Book struct {
	BookID           int64     `json:"book_id"`
	BookName         string    `json:"book_name"`
	ISBN             string    `json:"isbn"`
	PublishYear      int       `json:"publish_year"`
	PublisherID      int32     `json:"publisher_id"`
	BookTypeID       int16     `json:"book_type_id"`
	BookTypeDetailID int32     `json:"book_type_detail_id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	CreateUserID     int64     `json:"create_user_id"`
	UpdateUserID     int64     `json:"update_user_id"`
}

type (
	CreateBookPayload struct {
		BookName         string `json:"book_name"`
		HasISBN          bool   `json:"has_isbn"`
		ISBN             string `json:"book_isbn"`
		PublishYear      int    `json:"publish_date"`
		PublisherID      int32  `json:"publisher_id"`
		BookTypeID       int16  `json:"book_type_id"`
		BookTypeDetailID int32  `json:"book_type_detail_id"`
	}

	UpdateBookPayload struct {
		BookName         string    `json:"book_name"`
		ISBN             string    `json:"isbn"`
		PublishYear      time.Time `json:"publish_date"`
		PublisherID      int32     `json:"publisher_id"`
		BookTypeID       int16     `json:"book_type_id"`
		BookTypeDetailID int32     `json:"book_type_detail_id"`
	}
)

func (payload *CreateBookPayload) IsValid() (bool, map[string]string) {

	v := NewValidator()
	// BookName
	v.MustBeNotEmpty("book", payload.BookName)
	v.MustBeLongerThan("book", payload.BookName, 2)

	// ISBN -Year
	if payload.HasISBN {
		// length
		v.MustBeNotEmpty("isbn", payload.ISBN)
		v.MustBeLongerThan("isbn", payload.ISBN, 9)
		v.MustBeValidYear("publish_year", payload.PublishYear)
	}
	return v.IsValid(), v.errors
}

func (d *Domain) CreateBook(payload CreateBookPayload) (*Book, error) {

	// Check that book isbn is not repeated
	bookExist, _ := d.DB.BookRepo.GetByISBN(payload.ISBN)
	if bookExist != nil {
		return nil, ErrBookIsAlreadyExist
	}

	// Check that book isbn is not repeated
	bookExist, _ = d.DB.BookRepo.GetByName(payload.BookName)
	if bookExist != nil {
		return nil, ErrBookIsAlreadyExist
	}

	if !payload.HasISBN {
		payload.ISBN = "-"
	}

	book := &Book{
		BookName:         payload.BookName,
		ISBN:             payload.ISBN,
		PublishYear:      payload.PublishYear,
		PublisherID:      payload.PublisherID,
		BookTypeID:       payload.BookTypeID,
		BookTypeDetailID: bookExist.BookTypeDetailID,
		CreateUserID:     1,
		UpdateUserID:     1,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	book, err := d.DB.BookRepo.Create(book)
	if err != nil {
		return nil, err
	}

	return book, nil

}
