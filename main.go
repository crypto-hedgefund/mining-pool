package main

import (
	"fmt"
	"log"

	"main.go/config"
	"main.go/redis"
	"main.go/stats"
)

func main() {
	// Get env/config vars
	config, err := config.Setup()
	if err != nil {
		log.Fatal(err)
	}

	// Get redis client instance
	redisClient := redis.NewRedisClient(config)

	// Check if redis server is up
	pong, err := redisClient.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("redis reply: %v", pong)

	// Start services
	// 		Stats
	// 		Work
	// 		Payout

	statService := stats.NewStatsService()
	statService.Start()

	_ = config

	// Prevent goroutines from terminating early
	quit := make(chan bool)
	<-quit
}
