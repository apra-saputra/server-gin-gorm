package initializers

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var RdsClient *redis.Client

var CtxRds = context.Background()

func ConnectRedis() {
	RdsClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
