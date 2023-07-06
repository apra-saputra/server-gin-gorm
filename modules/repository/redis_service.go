package repository

import (
	"restapi/initializers"
	"restapi/modules/request"
	"restapi/modules/response"
	"time"
)

func (repo *RedisRepository) StoreValuesRedis(keysAndValues []request.StoreRedis) error {
	redisClient := repo.Redis
	ctx := initializers.CtxRds

	for _, item := range keysAndValues {
		err := redisClient.Set(ctx, item.Key, item.Value, 15*time.Minute).Err()

		if err != nil {
			return err
		}
	}
	return nil
}

func (repo *RedisRepository) ReadValueRedis(keys []string) ([]response.RedisResponse, error) {
	redisClient := repo.Redis
	ctx := initializers.CtxRds

	var result []response.RedisResponse

	for _, item := range keys {
		data, err := redisClient.Get(ctx, item).Result()
		if err != nil {
			return nil, err
		}

		store := response.RedisResponse{
			Key:        item,
			CacheValue: data,
		}

		result = append(result, store)
	}

	return result, nil
}
