package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/voltgizerz/go-redis/config"
	"github.com/voltgizerz/go-redis/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("===Start Redis Go Client===")

	// init redis
	redis := config.InitRedis()

	bookService := service.NewBookService(redis)

	bookService.ReadBook()
}
