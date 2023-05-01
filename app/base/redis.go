package base

import (
	"context"
	"github.com/go-redis/redis/v9"
	"os"
)

func InitRedis() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		DB:       0,
		Username: os.Getenv("REDIS_USERNAME"),
		Password: os.Getenv("REDIS_PASSWORD"),
	})

	var ctx = context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil
	}

	return client
}
