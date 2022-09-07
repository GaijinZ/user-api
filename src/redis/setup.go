package redis

import (
	"github.com/go-redis/redis/v8"
)

func RedisSetup() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.33.2:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client
}
