package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisService struct {
	client *redis.Client
}

// lux:cons prac
func NewService() (*RedisService, error) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	return &RedisService{
		client: client,
	}, nil
}

func (r *RedisService) Increment(ctx context.Context) (int64, error) {
	v, err := r.client.Incr(ctx, "count").Result()
	if err != nil {
		return 0, fmt.Errorf("failed to increment count: %w", err)
	}

	return v, nil
}
