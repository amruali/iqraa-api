package validator

import (
	"errors"
	"fmt"
)

var (
	ErrNoResult                         = errors.New("no result")
	ErrPasswordDoesotMatch              = errors.New("password doesnot match")
	ErrUserWithEmailAlreadyAlreadyExist = errors.New("user with email already exist")
	ErrUserWithUsernameAlreadyExist     = errors.New("user with username already exist")
	ErrInvalidCredentials               = errors.New("username or password are not correct")
	ErrEmailIsNotValid                  = errors.New("email is not valid")
	ErrBookIsAlreadyExist               = errors.New("book is already exist")
	ErrBookIsNotFound                   = errors.New("book is not found")
	ErrDataIsNotCashed                  = errors.New("these data is not cashed or ttl-ended")
)

type (
	ErrNotLongEnough struct {
		field  string
		length int
	}
	ErrIsRequired struct {
		field string
	}
)

func (e ErrNotLongEnough) Error() string {
	return fmt.Sprintf("%v should be at least %v", e.field, e.length)
}

func (e ErrIsRequired) Error() string {
	return fmt.Sprintf("%v is required", e.field)
}
