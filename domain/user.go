package domain

import (
	"time"
)

type User struct {
	Id             int64     `json:"user_id"`
	UserName       string    `json:"username"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"-"`
	UserTypeID     int8      `json:"user_type_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (d *Domain) GetUserByUserID(id int64) (*User, error) {
	user, err := d.DB.UserRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
