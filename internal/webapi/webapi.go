package webapi

import "github.com/vgekko/ani-go/internal/entity"

type WebAPI struct {
	Kodik
}

func NewWebAPI() *WebAPI {
	return &WebAPI{
		Kodik: NewKodikWebAPI(),
	}
}

type Kodik interface {
	SearchTitles(option, value string) (entity.KodikAPI, error)
}
