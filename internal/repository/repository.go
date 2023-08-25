package repository

import (
	"github.com/allegro/bigcache"
	"github.com/vgekko/anistream-content/internal/entity"
)

type Repository struct {
	CacheRepository
}

type CacheRepository interface {
	Get(key string) (entity.Content, error)
	Set(key string, data entity.Content) error
}

func NewRepository(cache *bigcache.BigCache) *Repository {
	return &Repository{CacheRepository: NewCacheRepository(cache)}
}
