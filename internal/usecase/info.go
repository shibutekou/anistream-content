package usecase

import (
	"context"
	"fmt"

	"github.com/vgekko/ani-go/internal/entity"
	"github.com/vgekko/ani-go/internal/repository/redis"
	"github.com/vgekko/ani-go/internal/webapi"
)

type InfoUseCase struct {
	kodik    webapi.Kodik
	infoRepo redis.Info
}

func NewInfoUseCase(kodik webapi.Kodik, infoRepo redis.Info) *InfoUseCase {
	return &InfoUseCase{
		kodik:    kodik,
		infoRepo: infoRepo,
	}
}

func (uc *InfoUseCase) Search(filter entity.TitleFilter) ([]entity.TitleInfo, error) {
	ctx := context.Background()



	var titleInfos = make([]entity.TitleInfo, 0)
	var err error

	// check the cache
	cacheKey := fmt.Sprintf("%s%s", filter.Option, filter.Value)

	// if data exists in cache, take it form there
	if exists := uc.infoRepo.Lookup(ctx, cacheKey); exists {
		titleInfos, err = uc.infoRepo.FromCache(ctx, cacheKey)
		if err != nil {
			return nil, fmt.Errorf("InfoUseCase.Search: %w", err)
		}

		return titleInfos, nil
	}

	// if data does not exists in cache
	results, err := uc.kodik.SearchTitles(filter.Option, filter.Value)
	if err != nil {
		return nil, fmt.Errorf("InfoUseCase.Search: %w", err)
	}

	titleInfos = toTitleInfo(results)

	// save data in cache
	err = uc.infoRepo.Cache(ctx, cacheKey, titleInfos)
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
