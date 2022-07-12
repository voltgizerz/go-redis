package main

import (
	"log"
	"os"

	"github.com/go-redis/redis"
)

type Redis struct {
	client *redis.Client
}

func InitRedis() Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
	})
	return Redis{client}
}

func (r Redis) Close() {
	r.client.Close()
}

func (r Redis) Ping() {
	pong, err := r.client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(pong, err)
}

func (r Redis) Set(key, value string) {
	err := r.client.Set(key, value, 0).Err()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully set key:", key, "value:", value)
}

func (r Redis) Get(key string) string {
	val, err := r.client.Get(key).Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Getting value for key:", key, "value:", val)
	return val
}