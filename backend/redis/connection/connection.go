package redis

import "github.com/go-redis/redis"

func NewConnection() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "your-redis-secret-password",
		DB:       0,
	})
}
