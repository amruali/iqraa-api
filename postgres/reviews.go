package postgres

import (
	"iqraa-api/domain"

	"github.com/go-pg/pg/v9"
)

type ReviewRepo struct {
	DB *pg.DB
}

func NewReviewRepo(DB *pg.DB) *ReviewRepo {
	return &ReviewRepo{DB: DB}
}

// Create
func (r *ReviewRepo) Create(review *domain.Review) (*domain.Review, error) {

	_, err := r.DB.Model(review).Returning("*").Insert()
	if err != nil {
		return nil, err
	}
	return review, nil

}

// GetByID
func (r *ReviewRepo) GetByID(reviewID int32) (*domain.Review, error) {
	review := &domain.Review{}
	err := r.DB.Model(review).
		Where("review_id = ?", reviewID).
		Select()
	if err != nil {
		return nil, err
	}
	return review, nil
}

// GetByBookID
func (r *ReviewRepo) GetByBookID(bookID int32) ([]domain.Review, error) {
	reviews := []domain.Review{}
	err := r.DB.Model(&reviews).
		Where("book_id = ?", bookID).
		Select()
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

// GetByBookName
func (r *ReviewRepo) GetByBookName(bookName string) ([]domain.Review, error) {
	reviews := []domain.Review{}
	err := r.DB.Model(&reviews).
		ColumnExpr("review.*").
		//ColumnExpr("a.id AS author__id, a.full_name AS author__name").
		Join("JOIN books b"). 
		JoinOn("b.book_id = review.book_id").
		JoinOn("a.name = ?", bookName).
		Select()
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

// GetByUserID
func (r *ReviewRepo) GetByUserID(userID int32) ([]domain.Review, error) {
	reviews := []domain.Review{}
	err := r.DB.Model(&reviews).
		Where("create_user_id = ?", userID).
		Select()
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

