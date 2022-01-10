package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
)

type RedisBooksRepo struct {
	RedisDB *redis.Client
}

func NewRedisBooksRepo(RedisDB *redis.Client) *RedisBooksRepo {
	return &RedisBooksRepo{RedisDB: RedisDB}
}

func (r *RedisBooksRepo) GetStrings(key string) (string, error) {
	booksString, err := r.RedisDB.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return booksString, err
}

func (r *RedisBooksRepo) SetStrings(key, value string) error {
	if err := r.RedisDB.Set(ctx, key, value, time.Duration(time.Hour*24*7)).Err(); err != nil {
		return err
	}
	return nil
}
