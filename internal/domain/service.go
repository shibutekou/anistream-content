package domain

import (
	"github.com/vgekko/ani-go/internal/model"
	"net/http"
)

type linkService interface {
	ByID(service, id string) (string, error)
	ByTitleName(title string) (string, error)
}

type infoService interface {
	ByID(service, id string) ([]model.TitleInfo, error)
	ByTitleName(title string) ([]model.TitleInfo, error)
}

type conversionService interface {
	toTitleInfo(api *model.KodikAPI) []model.TitleInfo
}

type Service struct {
	Link linkService
	Info infoService
}

func NewService(token string, client http.Client) *Service {
	linkServiceImpl := NewLinkService(token, client)
	conversionServiceImpl := NewConversionService()
	infoServiceImpl := NewInfoService(token, client, conversionServiceImpl)

	return &Service{
		Link: linkServiceImpl,
		Info: infoServiceImpl,
	}
}
