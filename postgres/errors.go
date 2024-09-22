package postgres

import (
	"errors"
)

var (
	ErrNoResult       = errors.New("no result")
	ErrBookIsNotFound = errors.New("book is not found")
)
