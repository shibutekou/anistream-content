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

var infoTmpl = `
<!doctype html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
<img src="https://i.kodik.biz/screenshots/seria/176463/1.jpg" width="400" height="300">
<iframe width="480" height="360" src={{.}}http://kodik.info/video/34204/89e3a2101a462afdb31f28d133a32880/720p
        allowfullscreen></iframe>
</body>
</html>
`

func (uc *InfoUseCase) ByKinopoiskID(id string) (entity.TitleInfos, error) {
	ctx := context.Background()

	var titleInfos entity.TitleInfos
	var err error

	// check the cache
	cacheKey := fmt.Sprintf("kinopoisk%s", id)

	// if data exists in cache, take it form there
	if exists := uc.redisRepo.Lookup(ctx, cacheKey); exists {
		titleInfos, err = uc.redisRepo.FromCache(ctx, cacheKey)
		if err != nil {
			return nil, fmt.Errorf("InfoUseCase.ByKinopoiskID: %w", err)
		}

		return titleInfos, nil
	}

	// if data does not exists in cache
	results, err := uc.webAPI.ResultsByKinopoiskID(id)
	if err != nil {
		return nil, fmt.Errorf("InfoUseCase.ByKinopoiskID.uc.WebAPI.ResultsByKinopoiskID: %w", err)
	}

	titleInfos = toTitleInfo(results)

	// save data in cache
	err = uc.redisRepo.Cache(ctx, cacheKey, titleInfos)
	if err != nil {
		return nil, fmt.Errorf("InfoUseCase.ByKinopoiskID.uc.redisRepo.Cache: %w", err)
	}

	return titleInfos, nil
}

func (uc *InfoUseCase) ByShikimoriID(id string) ([]entity.TitleInfo, error) {
	results, err := uc.webAPI.ResultsByShikimoriID(id)
	if err != nil {
		return nil, fmt.Errorf("InfoUseCase.ByShikimoriID.uc.WebAPI.ResultsByShikimoriID: %w", err)
	}

	titleInfos := toTitleInfo(results)

	return titleInfos, nil
}

func (uc *InfoUseCase) ByIMDbID(id string) ([]entity.TitleInfo, error) {
	results, err := uc.webAPI.ResultsByIMDbID(id)
	if err != nil {
		return nil, fmt.Errorf("InfoUseCase: ByIMDbID: uc.WebAPI.ResultsByIMDbID: %w", err)
	}

	titleInfos := toTitleInfo(results)

	return titleInfos, nil
}

func (uc *InfoUseCase) ByTitleName(title string) ([]entity.TitleInfo, error) {
	results, err := uc.webAPI.ResultsByTitle(title)
	if err != nil {
		return nil, fmt.Errorf("InfoUseCase: ByTitleName: uc.WebAPI.ResultsByTitle: %w", err)
	}

	titleInfos := toTitleInfo(results)

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

