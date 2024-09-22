package postgres

import (
	"context"
	"fmt"
	"iqraa-api/models"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	_ "github.com/lib/pq"
)

func ConnectPostgres(DB_URI string) *pg.DB {
	opt, err := pg.ParseURL(DB_URI)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	DB := pg.Connect(opt)
	return DB
}

func TestDBConnection(db *pg.DB) error {

	ctx := context.Background()

	// Run a simple query to check the connection

	_, err := db.ExecContext(ctx, "SELECT 1")

	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	// Create schema if not exists
	// _, err = db.Exec(`CREATE SCHEMA IF NOT EXISTS iqraa`)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }

	fmt.Println("Database connection successful!")

	// var tables []struct {
	// 	TableName string `pg:"tablename"` // Maps the column to the field
	// }

	// _, err = db.QueryContext(ctx, &tables, `
	// SELECT tablename
	// FROM pg_catalog.pg_tables
	// WHERE schemaname = 'iqraa'`)
	// if err != nil {

	// 	fmt.Println(err)
	// 	return fmt.Errorf("failed to list tables in schema %s: %w", "iqraa", err)
	// }

	// if len(tables) == 0 {
	// 	fmt.Printf("No tables found in schema: %s\n", "iqraa")
	// 	return nil
	// }

	// fmt.Printf("Tables in schema '%s':\n", "iqraa")

	// for _, table := range tables {
	// 	fmt.Println(table.TableName)
	// }

	_, err = db.Exec(`SET search_path TO iqraa`)

	if err != nil {
		fmt.Println(err)
		return err
	}

	// var authors []string

	// _, err = db.QueryContext(ctx, &authors, ` SELECT fname FROM authors`)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return fmt.Errorf("failed to list tables in schema %s: %w", "iqraa", err)
	// }

	// fmt.Println(authors)

	return nil
}

// Recreate the schema using go-pg ORM
func RecreateSchema(db *pg.DB) error {
	// Drop tables if they exist

	// Set search path to iqraa
	_, err := db.Exec(`SET search_path TO iqraa`)
	if err != nil {

		fmt.Println(err)
		return err
	}

	models := []interface{}{
		(*models.User)(nil),
		(*models.Author)(nil),
		(*models.Book)(nil),
		(*models.Quote)(nil),
		(*models.Review)(nil),
	}

	for _, model := range models {
		err := db.Model(model).DropTable(&orm.DropTableOptions{
			IfExists: false,
			Cascade:  true,
		})
		if err != nil {
			fmt.Println("Error DROP table:", err)
			return err
		}
	}

	// Create tables based on the model structs
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: false,
		})
		if err != nil {
			fmt.Println("Error CREATE table:", err)
			return err
		}
	}

	return nil
}
