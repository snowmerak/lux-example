package counter

import (
	"context"
	"prac/lib/client/redis"
)

type CounterService struct {
	redisService *redis.RedisService
}

// lux:cons prac
func NewService(redisService *redis.RedisService) *CounterService {
	return &CounterService{
		redisService: redisService,
	}
}

func (c *CounterService) Count() int64 {
	v, _ := c.redisService.Increment(context.TODO())
	return v
}
