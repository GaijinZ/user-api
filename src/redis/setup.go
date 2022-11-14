package redis

import (
	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func RedisSetup() {

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "redis",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
