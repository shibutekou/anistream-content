package usecase

import (
	"github.com/vgekko/anistream-content/internal/entity"
	"github.com/vgekko/anistream-content/internal/repository"
	"github.com/vgekko/anistream-content/internal/webapi"
)

type UseCase struct {
	ContentUseCase
}

func NewUseCase(cache repository.CacheRepository, kodik *webapi.WebAPI) *UseCase {
	return &UseCase{ContentUseCase: NewContentUseCase(kodik, cache)}
}

type ContentUseCase interface {
	Search(filter entity.TitleFilter) ([]entity.TitleContent, error)
}
