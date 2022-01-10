package domain

import (
	"encoding/json"
)

func (d *Domain) GetCashedStrings(key string, dataType interface{}) (interface{}, error) {
	cashedBookString, err := d.RedisDB.RedisBooksRepo.GetStrings(key)
	if err != nil {
		return nil, ErrDataIsNotCashed
	}

	data := GetCustomizedType(cashedBookString, dataType)

	return data, nil
}

func (d *Domain) SetCashedStrings(key, value string) error {
	err := d.RedisDB.RedisBooksRepo.SetStrings(key, value)
	if err != nil {
		return err
	}
	return nil
}

func GetCustomizedType(key string, dataType interface{}) interface{} {
	b := []byte(key)
	_ = json.Unmarshal(b, &dataType)
	return dataType
}
