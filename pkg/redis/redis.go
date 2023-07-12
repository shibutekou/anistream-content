package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/vgekko/ani-go/config"
	"golang.org/x/exp/slog"
)

const dbCache = 0

func NewClient(cfg config.Redis) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:                  cfg.Addr,
		Password:              "",
		DB:                    dbCache,
	})


	red := slog.String("redis", "connected")
	dat := slog.Int("database", dbCache)
	slog.LogAttrs(context.Background(), slog.LevelInfo, "redis", red, dat)


	return rdb
}
