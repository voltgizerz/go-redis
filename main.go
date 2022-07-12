package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("===Start Redis Go Client===")

	// init redis
	redis := InitRedis()
	defer redis.Close()

	example := "felix_key_example"
	redis.Set(example, "felix-test-value-this-is-a-value")

	data := redis.Get(example)
	log.Println("Data fetch from redis:", data)
}
