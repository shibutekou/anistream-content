package usecase

import (
	"context"
	"github.com/vgekko/ani-go/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	Link interface {
		ByKinopoiskID(id string) (string, error)
		ByShikimoriID(id string) (string, error)
		ByIMDbID(id string) (string, error)
		ByTitleName(title string) (string, error)
	}

	Info interface {
		ByKinopoiskID(id string) (entity.TitleInfos, error)
		ByShikimoriID(id string) ([]entity.TitleInfo, error)
		ByIMDbID(id string) ([]entity.TitleInfo, error)
		ByTitleName(title string) ([]entity.TitleInfo, error)
	}

	InfoRedisRepo interface {
		Lookup(ctx context.Context, key string) bool
		FromCache(ctx context.Context, key string) (entity.TitleInfos, error)
		Cache(ctx context.Context, key string, value entity.TitleInfos) error
	}

	KodikWebAPI interface {
		ResultsByKinopoiskID(id string) (entity.KodikAPI, error)
		ResultsByShikimoriID(id string) (entity.KodikAPI, error)
		ResultsByIMDbID(id string) (entity.KodikAPI, error)
		ResultsByTitle(title string) (entity.KodikAPI, error)
	}
)
