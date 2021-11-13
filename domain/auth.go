package domain

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type RegisterPayload struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Username        string `json:"username"`
}

func (payload *RegisterPayload) IsValid() (bool, map[string]string) {
	v := NewValidator()

	v.MustBeNotEmpty("email", payload.Email)
	v.MustBeValidEmail("email", payload.Email)

	v.MustBeNotEmpty("username", payload.Username)
	v.MustBeLongerThan("username", payload.Username, 4)

	v.MustBeNotEmpty("password", payload.Password)
	v.MustBeLongerThan("password", payload.Password, 8)

	v.MustBeNotEmpty("confirm_password", payload.ConfirmPassword)
	v.MustMatchPasswordAndConfrimPassword("password", payload.Password, payload.ConfirmPassword)

	return v.IsValid(), v.errors
}

func (d *Domain) Register(payload RegisterPayload) (*User, error) {

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

	user, err := d.DB.UserRepo.Create(payload.Email, payload.Username, hashedPassword)
	if err != nil {
		return nil, err
	}
	return user, nil
}

type LoginPayload struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *LoginPayload) IsValid() (bool, map[string]string) {
	v := NewValidator()

	//v.MustBeNotEmpty("username", l.UserName)

	v.MustBeNotEmpty("email", l.Email)
	v.MustBeValidEmail("email", l.Email)

	v.MustBeNotEmpty("password", l.Password)

	return v.IsValid(), v.errors
}

func (d *Domain) Login(payload LoginPayload) (*User, error) {
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
