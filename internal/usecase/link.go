package usecase

import (
	"fmt"
	"github.com/vgekko/ani-go/internal/entity"

	"github.com/vgekko/ani-go/internal/webapi"
	"github.com/vgekko/ani-go/pkg/normalize"
)

type LinkUseCase struct {
	kodik webapi.Kodik
}

func NewLinkUseCase(kodik webapi.Kodik) *LinkUseCase {
	return &LinkUseCase{kodik: kodik}
}

func (l *LinkUseCase) Search(filter entity.TitleFilter) (entity.Link, error) {
	results, err := l.kodik.SearchTitles(filter.Option, filter.Value)
	if err != nil {
		return entity.Link{}, fmt.Errorf("LinkUseCase.Search: %w", err)
	}

	url := normalize.URL(results.Results[0].Link)

	link := entity.Link{URL: url}

	return link, nil
}
