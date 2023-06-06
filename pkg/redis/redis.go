package redis

import (
	"github.com/redis/go-redis/v9"
)

const cacheDB = iota

type Redis struct {
	RedisClient *redis.Client
}

func New() *Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr:                  "localhost:6379",
		Password:              "",
		DB:                    cacheDB,
	})

	return &Redis{RedisClient: rdb}
}

func (r *Redis) Close() {
	if r.RedisClient != nil {
		r.RedisClient.Close()
	}
}
