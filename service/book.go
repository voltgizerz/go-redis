package service

import (
	"crypto/rand"
	log "github.com/sirupsen/logrus"
	"time"
	"unsafe"

	"github.com/voltgizerz/go-redis/config"
)

const REDIS_BOOK_KEY = "book_felix_key"

type BookInterface interface {
	ReadBook()
}

type Book struct {
	Redis config.RedisInterface
}

func NewBookService(redis config.RedisInterface) BookInterface {
	return Book{
		Redis: redis,
	}
}

func (b Book) ReadBook() {
	data := b.Redis.Get(REDIS_BOOK_KEY)
	if data != "" {
		log.Println("Data fetch from redis:", data)
		return
	} 
	time.Sleep(1 * time.Minute)
	b.Redis.Set(REDIS_BOOK_KEY, generate(256))
	log.Println("Finished Read Book!")
}

var alphabet = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generate(size int) string {
	b := make([]byte, size)
	rand.Read(b)
	for i := 0; i < size; i++ {
		b[i] = alphabet[b[i]%byte(len(alphabet))]
	}
	return *(*string)(unsafe.Pointer(&b))
}
