package postgres

import (
	"errors"
	"iqraa-api/models"

	"github.com/go-pg/pg/v10"
)

type AuthorRepo struct {
	DB *pg.DB
}

func NewAuthorRepo(DB *pg.DB) *AuthorRepo {
	return &AuthorRepo{DB: DB}
}

// Get Author By Name
// Get Author By ID
// Add Author
// Delete Author
// Update Author

// Get Authors By Nationality
// Get Authors By Books Type ID, Type Detail ID
// Get Top #N Authors By their Dowloadable Books IN Book Type ID, Type Detail
// Get Top #N Authors By their Dowloadable Books

func (a *AuthorRepo) Create(author *models.Author) (*models.Author, error) {
	_, err := a.DB.Model(author).Returning("*").Insert()
	if err != nil {
		return nil, err
	}
	return author, nil
}

func (a *AuthorRepo) GetByName(AuthorName string) (*models.Author, error) {
	author := &models.Author{}
	err := a.DB.Model(author).Where("full_name = ?", AuthorName).First()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, ErrNoResult
		}
		return nil, err
	}

	return author, nil
}
