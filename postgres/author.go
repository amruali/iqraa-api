package postgres

import "github.com/go-pg/pg/v9"

type AuthorRepo struct {
	DB *pg.DB
}

func NewAuthorRepo(DB *pg.DB) *AuthorRepo {
	return &AuthorRepo{DB : DB}
}