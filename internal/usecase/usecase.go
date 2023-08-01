package usecase

import (
	"github.com/vgekko/ani-go/internal/entity"
	"github.com/vgekko/ani-go/internal/repository/redis"
	"github.com/vgekko/ani-go/internal/webapi"
)

type UseCase struct {
	LinkUseCase
	InfoUseCase
}

func NewUseCase(redis *redis.RepositoryRedis, kodik *webapi.WebAPI) *UseCase {
	return &UseCase{
		LinkUseCase: NewLinkUseCase(kodik),
		InfoUseCase: NewInfoUseCase(kodik, redis.InfoRepository),
	}
}

type LinkUseCase interface {
	Search(filter entity.TitleFilter) (entity.Link, error)
}

type InfoUseCase interface {
	Search(filter entity.TitleFilter) ([]entity.TitleInfo, error)
}
