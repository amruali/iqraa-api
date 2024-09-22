package domain

import "iqraa-api/models"

func (d *Domain) GetTopDownloadedBooks() ([]models.Book, error) {
	// domain books count = 5
	books, err := d.DB.StatisticsRepo.GetByTopDownloaded(5)
	if err != nil {
		return nil, err
	}
	return books, nil
}
