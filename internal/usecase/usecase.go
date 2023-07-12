package usecase

import (
	"github.com/vgekko/ani-go/internal/entity"
	"github.com/vgekko/ani-go/internal/repository/postgres"
	"github.com/vgekko/ani-go/internal/repository/redis"
	"github.com/vgekko/ani-go/internal/webapi"
)

type UseCase struct {
	Link
	Info
	Auth
}

func NewUseCase(redis *redis.RepositoryRedis, postgres *postgres.RepositoryPostgres, kodik *webapi.WebAPI) *UseCase {
	return &UseCase{
		Link: NewLinkUseCase(kodik),
		Info: NewInfoUseCase(kodik, redis.Info),
		Auth: NewAuthUseCase(postgres.Auth),
	}
}

type Link interface {
	Search(option, value string) (string, error)
}

type Info interface {
	Search(option, value string) (entity.TitleInfos, error)
}

type Auth interface {
	SignUp()
}