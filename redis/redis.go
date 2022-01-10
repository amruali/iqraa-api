package redis

import (
	"log"

	"github.com/go-redis/redis/v8"
)

/*
func ConnectRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:   "localhost:6379",
		Password: "",   // no password set
		DB:       0,   // use default DB

	})
	return rdb
}
*/

func ConnectRedis(DB_URI string) *redis.Client {
	opt, err := redis.ParseURL(DB_URI)
	if err != nil {
		log.Fatalf("failed to connect to redis because %v", err)
	}
	rdb := redis.NewClient(opt)
	return rdb
}
