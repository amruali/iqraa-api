package postgres

import (
	"errors"
	"iqraa-api/models"

	"github.com/go-pg/pg/v10"
)

type StatisticsRepo struct {
	DB *pg.DB
}

func NewStatisticsRepo(DB *pg.DB) *StatisticsRepo {
	return &StatisticsRepo{DB: DB}
}

// Get Top sorted descending Downloadable #N Books
func (s *StatisticsRepo) GetByTopDownloaded(count int) ([]models.Book, error) {
	if count == -1 {
		count = 5
	}
	books := []models.Book{}
	err := s.DB.Model(&books).
		Order("downloads_number DESC").
		Limit(count).
		Select()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, ErrNoResult
		}
		return nil, err
	}
	return books, nil
}

// Get Highest Reviewed Books

// Get #N top Newest Books

// Get #N Most Downloadable Book for Specific Author

// Get #N top Highest reviewed Books

// Get #N top Books Based ON Specific User Prefer

// Get #N top Books These Days
