package usecase

import (
	"github.com/vgekko/anistream/internal/entity"
	"github.com/vgekko/anistream/internal/repository/redis"
	"github.com/vgekko/anistream/internal/webapi"
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
