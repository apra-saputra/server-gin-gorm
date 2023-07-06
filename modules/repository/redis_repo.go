package repository

import (
	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	Redis *redis.Client
}

func NewRedisRepository(redis *redis.Client) *RedisRepository {
	return (&RedisRepository{
		Redis: redis,
	})
}
