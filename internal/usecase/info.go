package usecase

import (
	"context"
	"fmt"
	"github.com/vgekko/anistream/internal/entity"
	"github.com/vgekko/anistream/internal/repository/redis"
	"github.com/vgekko/anistream/internal/webapi"
)

type InfoUseCaseImpl struct {
	kodik          webapi.Kodik
	infoRepository redis.InfoRepository
}

func NewInfoUseCase(kodik webapi.Kodik, infoRepository redis.InfoRepository) *InfoUseCaseImpl {
	return &InfoUseCaseImpl{
		kodik:          kodik,
		infoRepository: infoRepository,
	}
}

func (uc *InfoUseCaseImpl) Search(filter entity.TitleFilter) ([]entity.TitleInfo, error) {
	op := "InfoUseCase.Search"
	ctx := context.Background()

	var titleInfos []entity.TitleInfo
	var err error

	// check the cache if cache database is available
	// if data exists in cache, take it from there
	redisAvailable := uc.infoRepository.Healthcheck(ctx)
	key := fmt.Sprintf("%s:%s", filter.Option, filter.Value)

	if redisAvailable {
		exists, err := uc.infoRepository.Lookup(ctx, key)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		if exists {
			titleInfos, err = uc.infoRepository.FromCache(ctx, key)
			if err != nil {
				return nil, fmt.Errorf("%s: %w", op, err)
			}

			return titleInfos, nil
		}
	}

	// if data does not exists in cache
	results, err := uc.kodik.SearchTitles(filter.Option, filter.Value)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	titleInfos = toTitleInfo(results)

	// save data in cache
	if redisAvailable {
		err = uc.infoRepository.Cache(ctx, key, titleInfos)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
	}

	return titleInfos, nil
}

func toTitleInfo(src entity.KodikAPI) []entity.TitleInfo {
	var ti entity.TitleInfo
	titleInfos := make([]entity.TitleInfo, 0, len(src.Results))

	for _, v := range src.Results {
		ti.Title = v.Title
		ti.TitleOrig = v.TitleOrig
		ti.OtherTitle = v.OtherTitle
		ti.Year = v.Year
		ti.KinopoiskID = v.KinopoiskID
		ti.ShikimoriID = v.ShikimoriID
		ti.IMDbID = v.IMDbID
		ti.Screenshots = v.Screenshots

		titleInfos = append(titleInfos, ti)
	}

	return titleInfos
}
