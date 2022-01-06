package domain

import (
	"encoding/json"
)

func (d *Domain) GetCashedTopDownloads() ([]Book, error) {
	cashedBookString, err := d.RedisDB.RedisBooksRepo.GetTopDownloads()
	if err != nil {
		return nil, ErrDataIsNotCashed
	}

	books := []Book{}
	b := []byte(cashedBookString)
	_ = json.Unmarshal(b, &books)

	return books, nil
}

func (d *Domain) SetCashedTopDownloads(stringifiedBooks string) error {
	err := d.RedisDB.RedisBooksRepo.SetTopDownloads(stringifiedBooks)
	if err != nil {
		return err
	}
	return nil
}
