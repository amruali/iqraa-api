package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	topDownloadsKey = "top_downloads"
	ctx             = context.Background()
)

type RedisBooksRepo struct {
	RedisDB *redis.Client
}

func NewRedisBooksRepo(RedisDB *redis.Client) *RedisBooksRepo {
	return &RedisBooksRepo{RedisDB: RedisDB}
}

func (r *RedisBooksRepo) GetTopDownloads() (string, error) {
	booksString, err := r.RedisDB.Get(ctx, topDownloadsKey).Result()
	if err != nil {
		return "", err
	}
	return booksString, err
}

func (r *RedisBooksRepo) SetTopDownloads(value string) error {
	if err := r.RedisDB.Set(ctx, topDownloadsKey, value, time.Duration(time.Hour*24*7)).Err(); err != nil {
		return err
	}
	return nil
}
