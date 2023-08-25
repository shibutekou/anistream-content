package repository

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/allegro/bigcache"
	"github.com/vgekko/anistream-content/internal/entity"
)

type CacheRepositoryImpl struct {
	cache *bigcache.BigCache
}

func NewCacheRepository(cache *bigcache.BigCache) *CacheRepositoryImpl {
	return &CacheRepositoryImpl{cache: cache}
}

func (r *CacheRepositoryImpl) Get(key string) (entity.Content, error) {
	op := "CacheRepository.Get"
	var content entity.Content

	val, err := r.cache.Get(key)
	if err != nil {
		if errors.Is(err, bigcache.ErrEntryNotFound) {
			return entity.Content{}, err
		}
		return entity.Content{}, fmt.Errorf("%s:%s", op, err)
	}

	b := bytes.NewReader(val)

	// decode []byte to []entity.TitleContent
	if err := gob.NewDecoder(b).Decode(&content); err != nil {
		return entity.Content{}, err
	}

	return content, nil
}

func (r *CacheRepositoryImpl) Set(key string, data entity.Content) error {
	op := "CacheRepository.Set"
	var b bytes.Buffer

	// encodes data to bytes buffer
	if err := gob.NewEncoder(&b).Encode(data); err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}

	if err := r.cache.Set(key, b.Bytes()); err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}

	return nil
}
