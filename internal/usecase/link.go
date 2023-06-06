package usecase

import (
	"fmt"
	"strings"
)

type LinkUseCase struct {
	webAPI KodikWebAPI
}

func NewLinkUseCase(w KodikWebAPI) *LinkUseCase {
	return &LinkUseCase{webAPI: w}
}

func (l *LinkUseCase) ByKinopoiskID(id string) (string, error) {
	results, err := l.webAPI.ResultsByKinopoiskID(id)
	if err != nil {
		return "", fmt.Errorf("LinkUseCase: ByKinopoiskID: uc.WebAPI.ResultsByKinopoiskID: %w", err)
	}

	link := results.Results[0].Link

	link = strings.TrimLeft(link, "/")

	resp := fmt.Sprintf("http://%s", link)
	//return results.Results[0].Link, nil
	return resp, nil
}

func (l *LinkUseCase) ByShikimoriID(id string) (string, error) {
	results, err := l.webAPI.ResultsByShikimoriID(id)
	if err != nil {
		return "", fmt.Errorf("LinkUseCase: ByShikimoriID: uc.WebAPI.ResultsByShikimoriID: %w", err)
	}

	return results.Results[0].Link, nil
}

func (l *LinkUseCase) ByIMDbID(id string) (string, error) {
	results, err := l.webAPI.ResultsByIMDbID(id)
	if err != nil {
		return "", fmt.Errorf("LinkUseCase: ByIMDbID: uc.WebAPI.ResultsByIMDbID: %w", err)
	}

	return results.Results[0].Link, nil
}

func (l *LinkUseCase) ByTitleName(title string) (string, error) {
	results, err := l.webAPI.ResultsByTitle(title)
	if err != nil {
		return "", fmt.Errorf("LinkUseCase: ByTitleName: uc.WebAPI.ResultsByTitle: %w", err)
	}

	return results.Results[0].Link, nil
}
