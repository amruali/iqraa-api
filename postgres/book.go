package postgres

import (
	"errors"
	"iqraa-api/domain"

	"github.com/go-pg/pg/v9"
)

type BookRepo struct {
	DB *pg.DB
}

func NewBookRepo(DB *pg.DB) *BookRepo {
	return &BookRepo{DB: DB}
}

func (b *BookRepo) Create(book *domain.Book) (*domain.Book, error) {
	_, err := b.DB.Model(book).Returning("*").Insert()
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (b *BookRepo) GetByBookName(BookName string) (*domain.Book, error) {
	book := &domain.Book{}
	err := b.DB.Model(book).Where("book_name = ?", BookName).First()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, domain.ErrNoResult
		}
		return nil, err
	}

	return book, nil
}

func (b *BookRepo) GetByISBN(isbn string) (*domain.Book, error) {
	book := &domain.Book{}
	err := b.DB.Model(book).Where("isbn = ?", isbn).First()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, domain.ErrNoResult
		}
		return nil, err
	}
	return book, nil
}
