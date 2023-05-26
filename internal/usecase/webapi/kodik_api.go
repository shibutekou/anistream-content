package webapi

import (
	"encoding/json"
	"fmt"
	"github.com/vgekko/ani-go/internal/entity"
	"github.com/vgekko/ani-go/pkg/apperror"
	"net/http"
	"os"
	"time"
)

type KodikWebAPI struct {
	token  string
	client http.Client
}

func New() *KodikWebAPI {
	token := os.Getenv("KODIK_TOKEN")
	client := http.Client{Timeout: time.Second * 3}

	return &KodikWebAPI{token: token, client: client}
}

func (k *KodikWebAPI) ResultsByKinopoiskID(id string) (entity.KodikAPI, error) {
	var kodikResponse entity.KodikAPI

	url := fmt.Sprintf("https://kodikapi.com/search?token=%s&kinopoisk_id=%s", k.token, id)

	resp, err := k.client.Get(url)
	if err != nil {
		return entity.KodikAPI{}, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&kodikResponse); err != nil {
		return entity.KodikAPI{}, err
	}

	if len(kodikResponse.Results) == 0 {
		return entity.KodikAPI{}, apperror.ErrTitleNotFound
	}

	return kodikResponse, nil
}

func (k *KodikWebAPI) ResultsByShikimoriID(id string) (entity.KodikAPI, error) {
	var kodikResponse entity.KodikAPI

	url := fmt.Sprintf("https://kodikapi.com/search?token=%s&shikimori_id=%s", k.token, id)

	resp, err := k.client.Get(url)
	if err != nil {
		return entity.KodikAPI{}, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&kodikResponse); err != nil {
		return entity.KodikAPI{}, err
	}

	if len(kodikResponse.Results) == 0 {
		return entity.KodikAPI{}, apperror.ErrTitleNotFound
	}

	return kodikResponse, nil
}

func (k *KodikWebAPI) ResultsByIMDbID(id string) (entity.KodikAPI, error) {
	var kodikResponse entity.KodikAPI

	url := fmt.Sprintf("https://kodikapi.com/search?token=%s&imdb_id=%s", k.token, id)

	resp, err := k.client.Get(url)
	if err != nil {
		return entity.KodikAPI{}, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&kodikResponse); err != nil {
		return entity.KodikAPI{}, err
	}

	if len(kodikResponse.Results) == 0 {
		return entity.KodikAPI{}, apperror.ErrTitleNotFound
	}

	return kodikResponse, nil
}

func (k *KodikWebAPI) ResultsByTitle(title string) (entity.KodikAPI, error) {
	var kodikResponse entity.KodikAPI

	url := fmt.Sprintf("https://kodikapi.com/search?token=%s&title=%s", k.token, title)

	resp, err := k.client.Get(url)
	if err != nil {
		return entity.KodikAPI{}, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&kodikResponse); err != nil {
		return entity.KodikAPI{}, err
	}

	if len(kodikResponse.Results) == 0 {
		return entity.KodikAPI{}, apperror.ErrTitleNotFound
	}

	return kodikResponse, nil
}
