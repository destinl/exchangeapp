package config

import (
	"exchangeapp/global"

	"github.com/go-redis/redis"
)

// var RedisClient *redis.Client

func initRedis() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := RedisClient.Ping().Result()

	if err!= nil {
		panic("Failed to connect to Redis")
	}

	global.RedisClient = RedisClient
}