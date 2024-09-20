package postgres

import (
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
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

// Recreate the schema using go-pg ORM
func recreateSchema(db *pg.DB) error {
    // Drop tables if they exist
    models := []interface{}{
        (*User)(nil),
        (*Product)(nil),
    }
    for _, model := range models {
        err := db.Model(model).DropTable(&orm.DropTableOptions{
            IfExists: false,
            Cascade:  true,
        })
        if err != nil {
            return err
        }
    }

    // Create tables based on the model structs
    for _, model := range models {
        err := db.Model(model).CreateTable(&orm.CreateTableOptions{
            Temp: false,
        })
        if err != nil {
            return err
        }
    }

    return nil
}
