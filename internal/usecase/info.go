package usecase

import (
	"fmt"
	"github.com/vgekko/ani-go/internal/entity"
)

type InfoUseCase struct {
	webAPI KodikWebAPI
}

func NewInfoUseCase(w KodikWebAPI) *InfoUseCase {
	return &InfoUseCase{
		webAPI: w,
	}
}

func (uc *InfoUseCase) ByKinopoiskID(id string) ([]entity.TitleInfo, error) {
	results, err := uc.webAPI.ResultsByKinopoiskID(id)
	if err != nil {
		return nil, fmt.Errorf("InfoUseCase: ByKinopoiskID: uc.WebAPI.ResultsByKinopoiskID: %w", err)
	}

	titleInfos := toTitleInfo(results)

	return titleInfos, nil
}

func (uc *InfoUseCase) ByShikimoriID(id string) ([]entity.TitleInfo, error) {
	results, err := uc.webAPI.ResultsByShikimoriID(id)
	if err != nil {
		return nil, fmt.Errorf("InfoUseCase: ByShikimoriID: uc.WebAPI.ResultsByShikimoriID: %w", err)
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

