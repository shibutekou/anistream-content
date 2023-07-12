package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/vgekko/ani-go/config"
	"github.com/vgekko/ani-go/internal/entity"
)

type RepositoryRedis struct {
	Info
}

func NewRepositoryRedis(redis *redis.Client, cfg config.Redis) *RepositoryRedis {
	return &RepositoryRedis{
		Info: NewInfoRepo(redis, cfg.InfoCacheTTL),
	}
}

type Info interface {
	Lookup(ctx context.Context, key string) bool
	FromCache(ctx context.Context, key string) (entity.TitleInfos, error)
	Cache(ctx context.Context, key string, value entity.TitleInfos) error
}
