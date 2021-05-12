package main

import (
	"log"

	"main.go/config"
)

func main() {
	config, err := config.Setup()
	if err != nil {
		log.Fatal(err)
	}

	_ = config
}
