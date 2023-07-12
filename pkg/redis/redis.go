package redis

import (
	"github.com/redis/go-redis/v9"
	"github.com/vgekko/ani-go/config"
	"golang.org/x/exp/slog"
)

const dbCache = 0

func NewClient(cfg config.Redis) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: "",
		DB:       dbCache,
	})

	slog.Info("connected to redis: ", slog.Attr{Key: "database", Value: slog.IntValue(dbCache)})
	return rdb
}
