package dtos

import (
	"iqraa-api/validator"
)

type (
	RegisterPayload struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
		Username        string `json:"username"`
		FirstName       string `json:"first_name"`
		LastName        string `json:"last_name"`
	}

	LoginPayload struct {
		UserName string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

func (payload *RegisterPayload) IsValid() (bool, map[string]string) {
	v := validator.NewValidator()

	v.MustBeNotEmpty("email", payload.Email)
	v.MustBeValidEmail("email", payload.Email)

	v.MustBeNotEmpty("username", payload.Username)
	v.MustBeLongerThan("username", payload.Username, 4)

	v.MustBeNotEmpty("first_name", payload.FirstName)
	v.MustBeLongerThan("first_name", payload.FirstName, 3)

	v.MustBeNotEmpty("last_name", payload.LastName)
	v.MustBeLongerThan("last_name", payload.LastName, 3)

	v.MustBeNotEmpty("password", payload.Password)
	v.MustBeLongerThan("password", payload.Password, 8)

	v.MustBeNotEmpty("confirm_password", payload.ConfirmPassword)
	v.MustMatchPasswordAndConfrimPassword("password", payload.Password, payload.ConfirmPassword)

	return v.IsValid(), v.Errors
}

func (l *LoginPayload) IsValid() (bool, map[string]string) {
	v := validator.NewValidator()

	//v.MustBeNotEmpty("username", l.UserName)

	v.MustBeNotEmpty("email", l.Email)
	v.MustBeValidEmail("email", l.Email)

	v.MustBeNotEmpty("password", l.Password)

	return v.IsValid(), v.Errors
}
