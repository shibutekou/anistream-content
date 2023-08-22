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

func (r *CacheRepositoryImpl) Get(key string) ([]entity.TitleContent, error) {
	op := "CacheRepository.Get"
	var titleInfos []entity.TitleContent

	val, err := r.cache.Get(key)
	if err != nil {
		if errors.Is(err, bigcache.ErrEntryNotFound) {
			return nil, err
		}
		return nil, fmt.Errorf("%s:%s", op, err)
	}

	b := bytes.NewReader(val)

	// decode []byte to []entity.TitleContent
	if err := gob.NewDecoder(b).Decode(&titleInfos); err != nil {
		return nil, err
	}

	return titleInfos, nil
}

func (r *CacheRepositoryImpl) Set(key string, data []entity.TitleContent) error {
	var b bytes.Buffer

	// encodes data to bytes buffer
	if err := gob.NewEncoder(&b).Encode(data); err != nil {
		return err
	}

	if err := r.cache.Set(key, b.Bytes()); err != nil {
		return err
	}

	return nil
}
