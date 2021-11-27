package postgres

import (
	"iqraa-api/domain"

	"github.com/go-pg/pg/v9"
)

type QuoteRepo struct {
	DB *pg.DB
}

func NewQuoteRepo(DB *pg.DB) *QuoteRepo {
	return &QuoteRepo{DB: DB}
}

func (q *QuoteRepo) Create(quote *domain.Quote) (*domain.Quote, error) {
	_, err := q.DB.Model(quote).
		Returning("*").
		Insert()
	if err != nil {
		return nil, err
	}
	return quote, nil
}

func (q *QuoteRepo) GetByID(quoteID uint32) (*domain.Quote, error) {
	quote := &domain.Quote{}
	err := q.DB.Model(quote).
		Where("id = ?", quoteID).
		Select()

	if err != nil {
		return nil, err
	}
	return quote, nil
}

func (q *QuoteRepo) GetByBookID(quoteID uint32) ([]domain.Quote, error) {
	quotes := []domain.Quote{}
	err := q.DB.Model(&quotes).
		Where("book_id = ?", quoteID).
		Select()

	if err != nil {
		return nil, err
	}
	return quotes, nil
}
