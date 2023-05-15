package domain

import (
	"github.com/vgekko/ani-go/internal/model"
)

type ConversionServiceImpl struct{}

func NewConversionService() *ConversionServiceImpl {
	return &ConversionServiceImpl{}
}

func (c *ConversionServiceImpl) toTitleInfo(kodik model.KodikAPI) []model.TitleInfo {
	var titleInfos []model.TitleInfo

	for _, v := range kodik.Results {
		info := model.TitleInfo{
			Title:         v.Title,
			TitleOrig:     v.TitleOrig,
			OtherTitle:    v.OtherTitle,
			Year:          v.Year,
			LastSeason:    v.LastSeason,
			LastEpisode:   v.LastEpisode,
			EpisodesCount: v.EpisodesCount,
			KinopoiskID:   v.KinopoiskID,
			ImdbID:        v.ImdbID,
			ShikimoriID:   v.ShikimoriID,
			Screenshots:   v.Screenshots,
		}

		titleInfos = append(titleInfos, info)
	}

	return titleInfos
}
