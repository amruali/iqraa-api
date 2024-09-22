package main

import (
	"fmt"
	"iqraa-api/domain"
	"iqraa-api/handlers"
	"iqraa-api/postgres"
	"iqraa-api/redis"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	// Load the .env file in the current directory
	godotenv.Load()

	// or

	//godotenv.Load(".env")

	// Connet to Postgres DB
	DB := postgres.ConnectPostgres(os.Getenv("DB_URI"))

	defer DB.Close()

	err := postgres.TestDBConnection(DB)
	if err != nil {
		fmt.Println(err)
	}

	/* in case of having no dump - this was for creating DB tables using models */
	// err = postgres.RecreateSchema(DB)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	domainDB := domain.DB{
		UserRepo:       postgres.NewUserRepo(DB),
		BookRepo:       postgres.NewBookRepo(DB),
		AuthorRepo:     postgres.NewAuthorRepo(DB),
		ReviewRepo:     postgres.NewReviewRepo(DB),
		QuoteRepo:      postgres.NewQuoteRepo(DB),
		StatisticsRepo: postgres.NewStatisticsRepo(DB),
	}

	// Connect to Redis DB
	redisDB := redis.ConnectRedis(os.Getenv("REDIS_DB_URI"))
	defer redisDB.Close()

	domainRedisDB := domain.RedisDB{
		RedisStringsRepo: redis.NewRedisStringsRepo(redisDB),
	}

	d := &domain.Domain{DB: domainDB, RedisDB: domainRedisDB}

	r := handlers.SetupRouter(d)

	err = http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r)
	if err != nil {
		log.Fatalf("cannot start server %v", err)
	}
}
