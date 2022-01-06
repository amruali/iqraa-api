package domain

import (
	"encoding/json"
)

func (d *Domain) GetCashedStrings(key string) ([]Book, error) {
	cashedBookString, err := d.RedisDB.RedisBooksRepo.GetStrings(key)
	if err != nil {
		return nil, ErrDataIsNotCashed
	}

	books := []Book{}
	b := []byte(cashedBookString)
	_ = json.Unmarshal(b, &books)

	return books, nil
}

func (d *Domain) SetCashedStrings(key, value string) error {
	err := d.RedisDB.RedisBooksRepo.SetStrings(key, value)
	if err != nil {
		return err
	}
	return nil
}
