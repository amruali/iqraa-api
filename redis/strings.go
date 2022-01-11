package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	ctx = context.Background()
)

type RedisStringsRepo struct {
	RedisDB *redis.Client
}

func NewRedisStringsRepo(RedisDB *redis.Client) *RedisStringsRepo {
	return &RedisStringsRepo{RedisDB: RedisDB}
}

func (r *RedisStringsRepo) GetStrings(key string) (string, error) {
	booksString, err := r.RedisDB.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return booksString, err
}

func (r *RedisStringsRepo) SetStrings(key, value string) error {
	if err := r.RedisDB.Set(ctx, key, value, time.Duration(time.Hour*24*7)).Err(); err != nil {
		return err
	}
	return nil
}
