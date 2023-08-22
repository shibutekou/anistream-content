package webapi

import "github.com/vgekko/anistream-content/internal/entity"

type WebAPI struct {
	Kodik
}

func NewWebAPI() *WebAPI {
	return &WebAPI{Kodik: NewKodikWebAPI()}
}

type Kodik interface {
	SearchTitles(option, value string) ([]entity.TitleInfo, error)
}
