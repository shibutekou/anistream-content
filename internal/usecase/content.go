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

func (uc *ContentUseCaseImpl) Search(filter entity.TitleFilter) (entity.Content, error) {
	op := "ContentUseCase.Search"

	var (
		content entity.Content
		titles  []entity.Title
	)

	var err error

	// check the cache if cache database is available
	// if data exists in cache, take it from there
	key := fmt.Sprintf("%s:%s", filter.Opt, filter.Val)
	exists := true

	content, err = uc.cache.Get(key)
	if err != nil {
		if errors.Is(err, bigcache.ErrEntryNotFound) {
			exists = false
		}
	}

	if exists {
		return content, nil
	} else {
		titles, err = uc.kodik.SearchTitles(filter.Opt, filter.Val)
		if err != nil {
			return entity.Content{}, fmt.Errorf("%s:%w", op, err)
		}

		content.Titles = titles
		content.Total = int32(len(titles))

		// saving data to cache
		if err := uc.cache.Set(key, content); err != nil {
			return entity.Content{}, fmt.Errorf("%s:%w", op, err)
		}

		return content, nil
	}

}
