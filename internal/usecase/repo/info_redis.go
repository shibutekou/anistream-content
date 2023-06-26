package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/vgekko/ani-go/internal/entity"
	"github.com/vgekko/ani-go/pkg/redis"
	"time"
)

const notExists = iota

type InfoRedisRepo struct {
	*redis.Redis
	ttl time.Duration
}

func NewInfoRedis(redis *redis.Redis, ttl time.Duration) *InfoRedisRepo {
	return &InfoRedisRepo{redis, ttl}
}

func (r *InfoRedisRepo) Cache(ctx context.Context, key string, value entity.TitleInfos) error {
	err := r.RedisClient.Set(ctx, key, value, r.ttl).Err()
	if err != nil {
		return fmt.Errorf("InfoRedisRepo.Cache: %w", err)
	}

	return nil
}

func (r *InfoRedisRepo) Lookup(ctx context.Context, key string) bool {
	data, _ := r.RedisClient.Exists(ctx, key).Result()

	if data == notExists {
		return false
	}

	return true
}

func (r *InfoRedisRepo) FromCache(ctx context.Context, key string) (entity.TitleInfos, error) {
	var ti entity.TitleInfos
	var asBytes []byte

	if err := r.RedisClient.Get(ctx, key).Scan(&asBytes); err != nil {
		return nil, fmt.Errorf("InfoRedisRepo.FromCache: %w", err)
	}

	if err := json.Unmarshal(asBytes, &ti); err != nil {
		return nil, fmt.Errorf("InfoRedisRepo.FromCache: %w", err)
	}

	return ti, nil
}
