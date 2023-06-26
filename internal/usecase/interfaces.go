package usecase

import (
	"context"
	"github.com/vgekko/ani-go/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mock.go -package=usecase_test

type (
	Link interface {
		Search(option, value string) (string, error)
	}

	Info interface {
		Search(option, value string) (entity.TitleInfos, error)
	}

	InfoRedisRepo interface {
		Lookup(ctx context.Context, key string) bool
		FromCache(ctx context.Context, key string) (entity.TitleInfos, error)
		Cache(ctx context.Context, key string, value entity.TitleInfos) error
	}

	KodikWebAPI interface {
		SearchTitles(option, value string) (entity.KodikAPI, error)
	}

	UserPostgresRepo interface {
		Create()
	}
)
