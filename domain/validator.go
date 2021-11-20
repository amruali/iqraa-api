package domain

import "net/mail"

type Validator struct {
	errors map[string]string
}

func NewValidator() *Validator {
	return &Validator{errors: make(map[string]string)}
}

func (v *Validator) IsValid() bool {
	return len(v.errors) == 0
}

func (v *Validator) MustBeLongerThan(field, value string, high int) bool {
	if _, ok := v.errors[field]; ok {
		return false
	}

	if value == "" {
		return true
	}

	if len(value) < high {
		v.errors[field] = ErrNotLongEnough{
			field:  field,
			length: high,
		}.Error()
		return false
	}

	return true
}

func (v *Validator) MustBeValidYear(field string, value int32) bool {
	if _, ok := v.errors[field]; ok {
		return false
	}

	if value < 0 || value > 2022 {
		v.errors[field] = ErrIsRequired{
			field: field,
		}.Error()
		return false
	}
	return true
}

func (v *Validator) MustBeNotEmpty(field, value string) bool {
	if _, ok := v.errors[field]; ok {
		return false
	}
	if value == "" {
		v.errors[field] = ErrIsRequired{
			field: field,
		}.Error()
		return false
	}
	return true
}

func (v *Validator) MustBeValidEmail(field, email string) bool {
	if _, ok := v.errors[field]; ok {
		return false
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		v.errors[field] = ErrEmailIsNotValid.Error()
		return false
	}
	return true
}

/*
func (v *Validator) MustBeValidPassword(password string) bool {

}
*/

func (v *Validator) MustMatchPasswordAndConfrimPassword(field, password, confirmedPassword string) bool {
	if _, ok := v.errors[field]; ok {
		return false
	}

	if password != confirmedPassword {
		v.errors[field] = ErrPasswordDoesotMatch.Error()
		return false
	}
	return true

}
