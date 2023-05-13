package domain

import (
	"encoding/json"
	"fmt"
	"github.com/vgekko/ani-go/internal/model"
	"github.com/vgekko/ani-go/pkg/apperror"
	"net/http"
)

type LinkServiceImpl struct {
	token  string
	client http.Client
}

func NewLinkService(token string, client http.Client) *LinkServiceImpl {
	return &LinkServiceImpl{
		token:  token,
		client: client,
	}
}

func (l *LinkServiceImpl) ByID(service, id string) (string, error) {
	var kodikResponse model.KodikAPI

	url := fmt.Sprintf("https://kodikapi.com/search?token=%s&%s=%s", l.token, service, id)

	resp, err := l.client.Get(url)
	if err != nil {
		return "", err
	}

	if err := json.NewDecoder(resp.Body).Decode(&kodikResponse); err != nil {
		return "", err
	}

	if len(kodikResponse.Results) == 0 {
		return "", apperror.ErrTitleNotFound
	}

	link := kodikResponse.Results[0].Link

	return link, nil
}

func (l *LinkServiceImpl) ByTitleName(title string) (string, error) {
	var kodikResponse model.KodikAPI

	url := fmt.Sprintf("https://kodikapi.com/search?token=%s&title=%s", l.token, title)

	resp, err := l.client.Get(url)
	if err != nil {
		return "", err
	}

	if err := json.NewDecoder(resp.Body).Decode(&kodikResponse); err != nil {
		return "", err
	}

	if len(kodikResponse.Results) == 0 {
		return "", apperror.ErrTitleNotFound
	}

	link := kodikResponse.Results[0].Link

	return link, nil
}
