package domain

import (
	"encoding/json"
	"fmt"
	"github.com/vgekko/ani-go/internal/model"
	"github.com/vgekko/ani-go/pkg/apperror"
	"net/http"
)

type Link struct {
	client http.Client
	token  string
}

func NewLink(client http.Client, token string) *Link {
	return &Link{client: client, token: token}
}

func (k *Link) ByID(service, id string) (string, error) {
	var kodikResponse model.KodikAPI

	url := fmt.Sprintf("https://kodikapi.com/search?token=%s&%s=%s", k.token, service, id)

	resp, err := k.client.Get(url)
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
