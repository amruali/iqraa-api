package postgres

import (
	"log"

	"github.com/go-pg/pg/v9"
	_ "github.com/lib/pq"
)

func New(opts *pg.Options) *pg.DB {
	db := pg.Connect(opts)
	if db == nil {
		log.Fatal("Cannot Connect to DB")
	}
	return db
}
