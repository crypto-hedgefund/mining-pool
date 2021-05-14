package redis

import (
	"fmt"

	"github.com/go-redis/redis"

	"main.go/config"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(cfg config.Config) *RedisClient {
	// Get address from config
	redisAddress := fmt.Sprintf("%v:%v", cfg.Redis.Address, cfg.Redis.Port)

	// Connection to redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "",
		DB:       0,
	})

	return &RedisClient{client}
}

func (r *RedisClient) Ping() (string, error) {
	return r.client.Ping().Result()
}
