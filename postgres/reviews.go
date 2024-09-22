package postgres

import (
	"iqraa-api/models"

	"github.com/go-pg/pg/v10"
)

type ReviewRepo struct {
	DB *pg.DB
}

func NewReviewRepo(DB *pg.DB) *ReviewRepo {
	return &ReviewRepo{DB: DB}
}

// Create
func (r *ReviewRepo) Create(review *models.Review) (*models.Review, error) {

	_, err := r.DB.Model(review).Returning("*").Insert()
	if err != nil {
		return nil, err
	}
	return review, nil

}

// GetByID
func (r *ReviewRepo) GetByID(reviewID int32) (*models.Review, error) {
	review := &models.Review{}
	err := r.DB.Model(review).
		Where("review_id = ?", reviewID).
		Select()
	if err != nil {
		return nil, err
	}
	return review, nil
}

// GetByBookID
func (r *ReviewRepo) GetByBookID(bookID int32) ([]models.Review, error) {
	reviews := []models.Review{}
	err := r.DB.Model(&reviews).
		Where("book_id = ?", bookID).
		Select()
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

// GetByBookName
func (r *ReviewRepo) GetByBookName(bookName string) ([]models.Review, error) {
	reviews := []models.Review{}
	err := r.DB.Model(&reviews).
		ColumnExpr("review.*").
		//ColumnExpr("a.id AS author__id, a.full_name AS author__name").
		Join("JOIN books b").
		JoinOn("b.book_id = review.book_id").
		JoinOn("b.book_name = ?", bookName).
		Select()
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

// GetByUserID
func (r *ReviewRepo) GetByUserID(userID int32) ([]models.Review, error) {
	reviews := []models.Review{}
	err := r.DB.Model(&reviews).
		Where("create_user_id = (?)", userID).
		Select()
	if err != nil {
		return nil, err
	}
	return reviews, nil
}
