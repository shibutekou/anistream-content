package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/vgekko/ani-go/config"
	"github.com/vgekko/ani-go/internal/entity"
)

type RepositoryRedis struct {
	InfoRepository
}

func NewRepositoryRedis(redis *redis.Client, cfg config.Redis) *RepositoryRedis {
	return &RepositoryRedis{
		InfoRepository: NewInfoRepository(redis, cfg.InfoCacheTTL),
	}
}

type InfoRepository interface {
	Lookup(ctx context.Context, key string) (bool, error)
	FromCache(ctx context.Context, key string) ([]entity.TitleInfo, error)
	Cache(ctx context.Context, key string, value []entity.TitleInfo) error
	Healthcheck(ctx context.Context) bool
}
