package postgres

import (
	"iqraa-api/models"

	"github.com/go-pg/pg/v10"
)

type QuoteRepo struct {
	DB *pg.DB
}

func NewQuoteRepo(DB *pg.DB) *QuoteRepo {
	return &QuoteRepo{DB: DB}
}

func (q *QuoteRepo) Create(quote *models.Quote) (*models.Quote, error) {
	_, err := q.DB.Model(quote).
		Returning("*").
		Insert()
	if err != nil {
		return nil, err
	}
	return quote, nil
}

func (q *QuoteRepo) GetByID(quoteID uint32) (*models.Quote, error) {
	quote := &models.Quote{}
	err := q.DB.Model(quote).
		Where("id = ?", quoteID).
		Select()

	if err != nil {
		return nil, err
	}
	return quote, nil
}

func (q *QuoteRepo) GetByBookID(quoteID uint32) ([]models.Quote, error) {
	quotes := []models.Quote{}
	err := q.DB.Model(&quotes).
		Where("book_id = ?", quoteID).
		Select()

	if err != nil {
		return nil, err
	}
	return quotes, nil
}
