package usecase

import (
	"errors"
	"fmt"
	"github.com/allegro/bigcache"
	"github.com/vgekko/anistream-content/internal/entity"
	"github.com/vgekko/anistream-content/internal/repository"
	"github.com/vgekko/anistream-content/internal/webapi"
)

type ContentUseCaseImpl struct {
	kodik webapi.Kodik
	cache repository.CacheRepository
}

func NewContentUseCase(kodik webapi.Kodik, cache repository.CacheRepository) *ContentUseCaseImpl {
	return &ContentUseCaseImpl{
		kodik: kodik,
		cache: cache,
	}
}

func (uc *ContentUseCaseImpl) Search(filter entity.TitleFilter) ([]entity.TitleContent, error) {
	op := "ContentUseCase.Search"

	var titleContents []entity.TitleContent
	var err error

	// check the cache if cache database is available
	// if data exists in cache, take it from there
	key := fmt.Sprintf("%s:%s", filter.Opt, filter.Val)
	exists := true

	content, err := uc.cache.Get(key)
	if err != nil {
		if errors.Is(err, bigcache.ErrEntryNotFound) {
			exists = false
		}
	}

	if exists {
		return content, nil
	} else {
		titleContents, err = uc.kodik.SearchTitles(filter.Opt, filter.Val)
		if err != nil {
			return nil, fmt.Errorf("%s:%w", op, err)
		}

		// saving data to cache
		if err := uc.cache.Set(key, titleContents); err != nil {
			return nil, fmt.Errorf("%s:%w", op, err)
		}

		return titleContents, nil
	}
}
