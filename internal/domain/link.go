package domain

import (
	"encoding/json"
	"fmt"
	"github.com/vgekko/ani-go/internal/model"
	"net/http"
)

type Kodik struct {
	client http.Client
	token  string
}

func NewKodik(client http.Client, token string) *Kodik {
	return &Kodik{client: client, token: token}
}

func (k *Kodik) ByID(service, id string) (string, error) {
	var kodikResponse model.KodikAPI

	url := fmt.Sprintf("https://kodikapi.com/search?token=%s&%s=%s", k.token, service, id)
	resp, err := k.client.Get(url)
	if err != nil {
		return "", fmt.Errorf("error while sending request to %s: %w", url, err)
	}

	if err := json.NewDecoder(resp.Body).Decode(&kodikResponse); err != nil {
		return "", fmt.Errorf("error while decoding data: %w", err)
	}

	link := kodikResponse.Results[0].Link

	return link, nil
}
