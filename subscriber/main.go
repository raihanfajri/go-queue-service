package main

import (
	"log"

	"../adaptors/sub"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	subscriber := sub.GetSubscriber()

	subscriber.Subscribe()
}
