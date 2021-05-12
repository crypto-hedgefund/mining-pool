package main

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"

	"main.go/config"
)

func main() {
	// Get env/config vars
	config, err := config.Setup()
	if err != nil {
		log.Fatal(err)
	}

	// Get address from config
	redisAddress := fmt.Sprintf("%v:%v", config.Redis.Address, config.Redis.Port)

	// Connection to redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "",
		DB:       0,
	})

	// Check if redis server is up
	pong, err := client.Ping().Result()

	fmt.Println(pong, err)

	_ = client
	_ = config
}
