package usecase

import (
	"fmt"
	"github.com/vgekko/anistream-content/internal/entity"

	"github.com/vgekko/anistream-content/internal/webapi"
	"github.com/vgekko/anistream-content/pkg/normalize"
)

type LinkUseCaseImpl struct {
	kodik webapi.Kodik
}

func NewLinkUseCase(kodik webapi.Kodik) *LinkUseCaseImpl {
	return &LinkUseCaseImpl{kodik: kodik}
}

func (l *LinkUseCaseImpl) Search(filter entity.TitleFilter) (entity.Link, error) {
	op := "LinkUseCase.Search"

	results, err := l.kodik.SearchTitles(filter.Option, filter.Value)
	if err != nil {
		return entity.Link{}, fmt.Errorf("%s: %w", op, err)
	}

	url := normalize.URL(results.Results[0].Link)

	link := entity.Link{URL: url}

	return link, nil
}
