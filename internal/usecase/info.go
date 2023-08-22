package usecase

import (
	"errors"
	"fmt"
	"github.com/allegro/bigcache"
	"github.com/vgekko/anistream-content/internal/entity"
	"github.com/vgekko/anistream-content/internal/repository"
	"github.com/vgekko/anistream-content/internal/webapi"
)

type InfoUseCaseImpl struct {
	kodik webapi.Kodik
	cache repository.CacheRepository
}

func NewInfoUseCase(kodik webapi.Kodik, cache repository.CacheRepository) *InfoUseCaseImpl {
	return &InfoUseCaseImpl{
		kodik: kodik,
		cache: cache,
	}
}

func (uc *InfoUseCaseImpl) Search(filter entity.TitleFilter) ([]entity.TitleInfo, error) {
	op := "InfoUseCase.Search"

	var titleInfos []entity.TitleInfo
	var err error

	// check the cache if cache database is available
	// if data exists in cache, take it from there
	key := fmt.Sprintf("%s:%s", filter.Opt, filter.Val)
	exists := true

	val, err := uc.cache.Get(key)
	if err != nil {
		if errors.As(err, &bigcache.ErrEntryNotFound) {
			exists = false
		}
	}

	if exists {
		return val, nil
	} else {
		titleInfos, err = uc.kodik.SearchTitles(filter.Opt, filter.Val)
		if err != nil {
			return nil, fmt.Errorf("%s:%w", op, err)
		}

		// saving data to cache
		if err := uc.cache.Set(key, titleInfos); err != nil {
			return nil, fmt.Errorf("%s:%w", op, err)
		}

		return titleInfos, nil
	}
}
