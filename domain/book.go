package domain

import (
	"fmt"
	"iqraa-api/dtos"
	"iqraa-api/models"
	"time"
)


func (d *Domain) CreateBook(payload dtos.CreateBookPayload) (*models.Book, error) {

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

	newbook := &models.Book{
		BookName:         payload.BookName,
		ISBN:             payload.ISBN,
		PublishYear:      2021,
		PublisherID:      1,
		BookTypeID:       1,
		BookTypeDetailID: 1,
		BookAuthorID:     1,
		CreateUserID:     1,
		UpdateUserID:     1,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	/*payload.BookName, payload.ISBN, payload.PublishYear*/
	book, err := d.DB.BookRepo.Create(newbook)
	if err != nil {
		fmt.Println("سيد جلال عبدالحليم")
		return nil, err
	}

	return book, nil
}

func (d *Domain) UpdateBook(payload dtos.UpdateBookPayload, bookID int64) error {
	newbook := &models.Book{
		ID:               bookID,
		BookName:         payload.BookName,
		ISBN:             payload.ISBN,
		PublishYear:      payload.PublishYear,
		PublisherID:      payload.PublisherID,
		BookTypeID:       payload.BookTypeID,
		BookTypeDetailID: payload.BookTypeDetailID,
		BookAuthorID:     payload.BookAuthorID,
		UpdateUserID:     1,
		UpdatedAt:        time.Now(),
		Brief:            payload.Brief,
	}

	err := d.DB.BookRepo.UpdateByID(newbook)
	if err != nil {
		fmt.Println(err)
		return ErrBookIsNotFound
	}
	return nil
}

func (d *Domain) GetBook(bookID int64) (*models.Book, error) {
	book, err := d.DB.BookRepo.GetByID(bookID)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (d *Domain) DeleteBook(bookID int64) error {
	err := d.DB.BookRepo.Delete(bookID)
	if err != nil {
		return err
	}
	return nil
}
