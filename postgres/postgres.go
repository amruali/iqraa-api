package postgres

import (
	"github.com/go-pg/pg/v9"
	_ "github.com/lib/pq"
)

func ConnectPostgres(DB_URI string) *pg.DB {
	opt, err := pg.ParseURL(DB_URI)
	if err != nil {
		panic(err)
	}
	DB := pg.Connect(opt)
	return DB
}
