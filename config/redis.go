package config

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/go-redis/redis"
)

type RedisInterface interface {
	Ping()
	Set(key, value string)
	Get(key string) string
	Close()
}

type Redis struct {
	client *redis.Client
}

func InitRedis() RedisInterface {
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})

	return Redis{
		client: client,
	}
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
	expireTime := 5 * time.Minute
	err := r.client.Set(key, value, expireTime).Err()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully set key:", key, "value:", value)
}

func (r Redis) Get(key string) string {
	val, err := r.client.Get(key).Result()
	if err != nil {
		log.Println("data not found on redis with key: " + key)
		return ""
	}
	log.Println("Data found on redis, getting value for key:", key, "value:", val)
	return val
}
