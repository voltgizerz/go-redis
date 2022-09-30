package service

import (
	"log"
	"time"

	"github.com/voltgizerz/go-redis/config"
)

const REDIS_BOOK_KEY = "book_felix_key"

type Book struct {
	Redis config.Redis
}

func NewBookService(redis config.Redis) Book {
	return Book{
		Redis: redis,
	}
}

func (b *Book) ReadBook() {
	data := b.Redis.Get(REDIS_BOOK_KEY)
	if data != "" {
		log.Println("Data fetch from redis:", data)
	}

	b.Redis.Set("WOWOS", REDIS_BOOK_KEY)
	time.Sleep(2 * time.Minute)
	log.Println("Data not found on redis :3")
}
