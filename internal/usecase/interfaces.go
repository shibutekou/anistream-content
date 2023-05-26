package usecase

import (
	"github.com/vgekko/ani-go/internal/entity"
)

type (
	Link interface {
		ByKinopoiskID(id string) (string, error)
		ByShikimoriID(id string) (string, error)
		ByIMDbID(id string) (string, error)
		ByTitleName(title string) (string, error)
	}

	Info interface {
		ByKinopoiskID(id string) ([]entity.TitleInfo, error)
		ByShikimoriID(id string) ([]entity.TitleInfo, error)
		ByIMDbID(id string) ([]entity.TitleInfo, error)
		ByTitleName(title string) ([]entity.TitleInfo, error)
	}

	KodikWebAPI interface {
		ResultsByKinopoiskID(id string) (entity.KodikAPI, error)
		ResultsByShikimoriID(id string) (entity.KodikAPI, error)
		ResultsByIMDbID(id string) (entity.KodikAPI, error)
		ResultsByTitle(title string) (entity.KodikAPI, error)
	}
)
