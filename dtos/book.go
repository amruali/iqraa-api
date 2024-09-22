package dtos

import (
	"iqraa-api/validator"
)

type (

	CreateBookPayload struct {
		BookName         string `json:"book_name"`
		HasISBN          bool   `json:"has_isbn"`
		ISBN             string `json:"book_isbn"`
		PublishYear      int32  `json:"publish_date"`
		PublisherID      int32  `json:"publisher_id"`
		BookTypeID       int32  `json:"book_type_id"`
		BookTypeDetailID int32  `json:"book_type_detail_id"`
	}

	UpdateBookPayload struct {
		BookName         string `json:"book_name"`
		HasISBN          bool   `json:"has_isbn"`
		ISBN             string `json:"isbn"`
		PublishYear      int32  `json:"publish_year"`
		PublisherID      int32  `json:"publisher_id"`
		BookTypeID       int32  `json:"book_type_id"`
		BookTypeDetailID int32  `json:"book_type_detail_id"`
		BookAuthorID     int32  `json:"book_author_id"`
		Brief            string `json:"brief"`
	}
)

func (payload *UpdateBookPayload) IsValid() (bool, map[string]string) {
	v := validator.NewValidator()

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
	return v.IsValid(), v.Errors
}



func (payload CreateBookPayload) IsValid() (bool, map[string]string) {

	v := validator.NewValidator()
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
	return v.IsValid(), v.Errors
}
