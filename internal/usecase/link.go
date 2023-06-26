package usecase

import (
	"fmt"
	"github.com/vgekko/ani-go/pkg/normalize"
)

type LinkUseCase struct {
	webAPI KodikWebAPI
}

func NewLinkUseCase(w KodikWebAPI) *LinkUseCase {
	return &LinkUseCase{webAPI: w}
}

func (l *LinkUseCase) Search(option, value string) (string, error) {
	results, err := l.webAPI.SearchTitles(option, value)
	if err != nil {
		return "", fmt.Errorf("LinkUseCase.Search: %w", err)
	}

	link := normalize.Link(results.Results[0].Link)
	return link, nil
}
