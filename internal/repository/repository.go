package repository

import (
	"github.com/allegro/bigcache"
	"github.com/vgekko/anistream-content/internal/entity"
)

type Repository struct {
	CacheRepository
}

type CacheRepository interface {
	Get(key string) ([]entity.TitleContent, error)
	Set(key string, data []entity.TitleContent) error
}

func NewRepository(cache *bigcache.BigCache) *Repository {
	return &Repository{CacheRepository: NewCacheRepository(cache)}
}
