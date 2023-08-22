package webapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/vgekko/anistream-content/internal/entity"
	"github.com/vgekko/anistream-content/pkg/apperror"
)

const baseURLSearch = "https://kodikapi.com/search?"

type KodikWebAPI struct {
	token  string
	client http.Client
}

func NewKodikWebAPI() *KodikWebAPI {
	token := os.Getenv("KODIK_TOKEN")
	client := http.Client{Timeout: time.Second * 3}

	return &KodikWebAPI{token: token, client: client}
}

func (k *KodikWebAPI) SearchTitles(option, value string) ([]entity.TitleInfo, error) {
	var kodikResponse entity.KodikAPI

	url := fmt.Sprintf("%stoken=%s&%s=%s", baseURLSearch, k.token, option, value)

	resp, err := k.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Kodik.SearchTitles: %w", err)
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&kodikResponse); err != nil {
		return nil, err
	}

	if kodikResponse.Total == 0 {
		return nil, apperror.ErrTitleNotFound
	}

	return toTitleInfo(kodikResponse), nil
}

func toTitleInfo(src entity.KodikAPI) []entity.TitleInfo {
	var ti entity.TitleInfo
	titleInfos := make([]entity.TitleInfo, 0, len(src.Results))

	for _, v := range src.Results {
		ti.Link = v.Link
		ti.Title = v.Title
		ti.TitleOrig = v.TitleOrig
		ti.OtherTitle = v.OtherTitle
		ti.Year = v.Year
		ti.KinopoiskID = v.KinopoiskID
		ti.ShikimoriID = v.ShikimoriID
		ti.IMDbID = v.IMDbID
		ti.Screenshots = v.Screenshots

		titleInfos = append(titleInfos, ti)
	}

	var uniqueTitles []entity.TitleInfo
	seen := make(map[string]bool)

	for _, v := range titleInfos {
		if !seen[v.Title] {
			seen[v.Title] = true
			uniqueTitles = append(uniqueTitles, v)
		}
	}

	return uniqueTitles
}
