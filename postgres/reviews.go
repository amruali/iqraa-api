package postgres

import (
	"github.com/go-pg/pg/v9"
)

type ReviewRepo struct {
	DB *pg.DB
}

func NewReviewRepo(DB *pg.DB) *AuthorRepo {
	return &AuthorRepo{DB: DB}
}

// Create

// GetByID

// GetByBookID

// GetByBookName

// GetByUserID

// Get Highest Reviewed Books