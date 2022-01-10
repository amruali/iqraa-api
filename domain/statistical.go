package domain

func (d *Domain) GetTopDownloadedBooks() ([]Book, error) {
	// domain books count = 5
	books, err := d.DB.StatisticsRepo.GetByTopDownloaded(5)
	if err != nil {
		return nil, err
	}
	return books, nil
}
