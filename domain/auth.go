package domain

import (
	"errors"
	"iqraa-api/dtos"
	"iqraa-api/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (d *Domain) Register(payload dtos.RegisterPayload) (*models.User, error) {

	// Check that email is not taken
	userExist, _ := d.DB.UserRepo.GetByEmail(payload.Email)
	if userExist != nil {
		return nil, ErrUserWithEmailAlreadyAlreadyExist
	}
	// check that username is not taken

	userExist, _ = d.DB.UserRepo.GetByUserName(payload.Username)
	if userExist != nil {
		return nil, ErrUserWithUsernameAlreadyExist
	}

	//  hash tha password
	hashedPassword, err := d.EncryptPassword(payload.Password)
	if err != nil {
		return nil, err
	}

	// create New User
	user := &models.User{
		Email:          payload.Email,
		Username:       payload.Username,
		HashedPassword: hashedPassword,
		UserTypeID:     3,
		FirstName:      payload.FirstName,
		LastName:       payload.LastName,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	registeredUser, err := d.DB.UserRepo.Create(user)
	if err != nil {
		return nil, err
	}
	return registeredUser, nil
}

func (d *Domain) Login(payload dtos.LoginPayload) (*models.User, error) {
	user, err := d.DB.UserRepo.GetByEmail(payload.Email)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(payload.Password))
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	return user, nil

}

func (d *Domain) EncryptPassword(plain string) (HashedPassword []byte, err error) {
	if len(plain) <= 7 {
		return HashedPassword, errors.New("password length should be greater than seven")
	}
	HashedPassword, err = bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return
}
