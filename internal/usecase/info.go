package usecase

import (
	"context"
	"fmt"
	"github.com/vgekko/ani-go/internal/entity"
)

type InfoUseCase struct {
	webAPI KodikWebAPI
	redisRepo InfoRedisRepo
}

func NewInfoUseCase(w KodikWebAPI, r InfoRedisRepo) *InfoUseCase {
	return &InfoUseCase{
		webAPI: w,
		redisRepo: r,
	}
}

func (uc *InfoUseCase) Search(option, value string) (entity.TitleInfos, error) {
	ctx := context.Background()

	var titleInfos entity.TitleInfos
	var err error

	// check the cache
	cacheKey := fmt.Sprintf("%s%s", option, value)

	// if data exists in cache, take it form there
	if exists := uc.redisRepo.Lookup(ctx, cacheKey); exists {
		titleInfos, err = uc.redisRepo.FromCache(ctx, cacheKey)
		if err != nil {
			return nil, fmt.Errorf("InfoUseCase.Search: %w", err)
		}

		return titleInfos, nil
	}

	// if data does not exists in cache
	results, err := uc.webAPI.SearchTitles(option, value)
	if err != nil {
		return nil, fmt.Errorf("InfoUseCase.Search: %w", err)
	}

	titleInfos = toTitleInfo(results)

	// save data in cache
	err = uc.redisRepo.Cache(ctx, cacheKey, titleInfos)
	if err != nil {
		return nil, fmt.Errorf("InfoUseCase.Search: %w", err)
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

