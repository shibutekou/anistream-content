package usecase

import (
	"fmt"

	"github.com/vgekko/ani-go/internal/webapi"
	"github.com/vgekko/ani-go/pkg/normalize"
)

type LinkUseCase struct {
	kodik webapi.Kodik
}

func NewLinkUseCase(kodik webapi.Kodik) *LinkUseCase {
	return &LinkUseCase{kodik: kodik}
}

func (l *LinkUseCase) Search(option, value string) (string, error) {
	results, err := l.kodik.SearchTitles(option, value)
	if err != nil {
		return "", fmt.Errorf("LinkUseCase.Search: %w", err)
	}

	link := normalize.Link(results.Results[0].Link)
	return link, nil
}
