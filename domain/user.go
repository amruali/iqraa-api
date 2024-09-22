package domain

import (
	"iqraa-api/dtos"
	"iqraa-api/models"
)

func (d *Domain) GetUserByUserID(id int64) (*models.User, error) {
	user, err := d.DB.UserRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (d *Domain) GetUserInfo(username string) (*dtos.Profile, error) {
	user, err := d.DB.UserRepo.GetByUserName(username)
	if err != nil {
		return nil, err
	}

	userProfile := &dtos.Profile{
		Username:   user.Email,
		Email:      user.Email,
		UserTypeID: user.UserTypeID,
		CreatedAt:  user.CreatedAt,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Image:      user.Image,
		Settings:   user.Settings,
	}
	
	return userProfile, nil
}
