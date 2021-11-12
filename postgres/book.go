package postgres

import "github.com/go-pg/pg/v9"

type BookRepo struct {
	DB *pg.DB
}

func NewBookRepo(DB *pg.DB) *BookRepo {
	return &BookRepo{DB: DB}
}
