package main

import (
	"fmt"
	"iqraa-api/domain"
	"iqraa-api/handlers"
	"iqraa-api/postgres"
	"log"
	"net/http"
	"os"

	"github.com/go-pg/pg/v9"
	"github.com/joho/godotenv"
)

func main() {


	// Load the .env file in the current directory
	godotenv.Load()

	// or

	//godotenv.Load(".env")

	opt, err := pg.ParseURL(os.Getenv("DB_URI"))
	if err != nil {
		panic(err)
	}

	DB := pg.Connect(opt)

	/*
		ctx := context.Background()

		if err := DB.Ping(ctx); err != nil {
			panic(err)
		}
	*/

	defer DB.Close()

	domainDB := domain.DB{
		UserRepo:       postgres.NewUserRepo(DB),
		BookRepo:       postgres.NewBookRepo(DB),
		AuthorRepo:     postgres.NewAuthorRepo(DB),
		ReviewRepo:     postgres.NewReviewRepo(DB),
		QuoteRepo:      postgres.NewQuoteRepo(DB),
		StatisticsRepo: postgres.NewStatisticsRepo(DB),
	}

	d := &domain.Domain{DB: domainDB}

	r := handlers.SetupRouter(d)

	err = http.ListenAndServe(fmt.Sprintf(":%s", /*"8080"*/ os.Getenv("PORT")), r)
	if err != nil {
		log.Fatalf("cannot start server %v", err)
	}
}
