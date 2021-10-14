package redisdb

import (
	"context"

	"github.com/IDarar/test-task/internal/config"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func NewRedisDB(cfg config.Config) (*redis.Client, error) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	status := rdb.Ping(ctx)
	if status.Err() != nil {
		return nil, status.Err()
	}

	return rdb, status.Err()
}
