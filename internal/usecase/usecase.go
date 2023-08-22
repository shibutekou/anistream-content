package usecase

import (
	"github.com/vgekko/anistream-content/internal/entity"
	"github.com/vgekko/anistream-content/internal/repository"
	"github.com/vgekko/anistream-content/internal/webapi"
)

type UseCase struct {
	InfoUseCase
}

func NewUseCase(cache repository.CacheRepository, kodik *webapi.WebAPI) *UseCase {
	return &UseCase{InfoUseCase: NewInfoUseCase(kodik, cache)}
}

type InfoUseCase interface {
	Search(filter entity.TitleFilter) ([]entity.TitleInfo, error)
}
