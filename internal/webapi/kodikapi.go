package webapi

import (
	"encoding/json"
	"fmt"
	"github.com/vgekko/anistream-content/pkg/apperror"
	"net/http"
	"os"
	"time"

	"github.com/vgekko/anistream-content/internal/entity"
)

const baseURLSearch = "https://kodikapi.com/search?with_material_data=true&"

// error response texts from kodikapi.com
const (
	kErrNoSearchParams        = "Не указан хотя бы один параметр для поиска"
	kErrMissingOrInvalidToken = "Отсутствует или неверный токен"
)

type KodikWebAPI struct {
	token  string
	client http.Client
}

func NewKodikWebAPI() *KodikWebAPI {
	token := os.Getenv("KODIK_TOKEN")
	client := http.Client{Timeout: time.Second * 3}

	return &KodikWebAPI{token: token, client: client}
}

// SearchTitles sends a request to external API and gets data about titles
func (k *KodikWebAPI) SearchTitles(option, value string) ([]entity.TitleContent, error) {
	var kodikResponse entity.KodikAPI

	url := fmt.Sprintf("%stoken=%s&%s=%s", baseURLSearch, k.token, option, value)

	resp, err := k.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		if err := json.NewDecoder(resp.Body).Decode(&kodikResponse); err != nil {
			return nil, err
		}

		if kodikResponse.Total == 0 {
			return nil, apperror.ErrTitleNotFound
		}

		return toTitleContent(kodikResponse), nil
	case http.StatusInternalServerError:
		var kErr entity.KodikError

		if err := json.NewDecoder(resp.Body).Decode(&kErr); err != nil {
			return nil, err
		}

		switch kErr.Error {
		case kErrNoSearchParams:
			return nil, apperror.ErrNoSearchParams
		case kErrMissingOrInvalidToken:
			return nil, apperror.ErrMissingOrInvalidToken
		default:
			return nil, apperror.ErrUnknown
		}
	default:
		return nil, err
	}
}

func toTitleContent(src entity.KodikAPI) []entity.TitleContent {
	var ti entity.TitleContent
	titleContents := make([]entity.TitleContent, 0, len(src.Results))

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
		ti.AnimeStatus = v.AnimeStatus
		ti.AnimeDescription = v.AnimeDescription
		ti.PosterURL = v.PosterURL
		ti.Duration = v.Duration
		ti.KinopoiskRating = v.KinopoiskRating
		ti.IMDbRating = v.IMDbRating
		ti.ShikimoriRating = v.ShikimoriRating
		ti.PremiereWorld = v.PremiereWorld
		ti.EpisodesTotal = v.EpisodesTotal
		ti.Writers = v.Writers

		titleContents = append(titleContents, ti)
	}

	return filterUnique(titleContents)
}

// filterUnique removes duplicate title contents from slice
func filterUnique(titleContents []entity.TitleContent) []entity.TitleContent {
	var uniqueTitles []entity.TitleContent
	seen := make(map[string]bool)

	for _, v := range titleContents {
		if !seen[v.Title] {
			seen[v.Title] = true
			uniqueTitles = append(uniqueTitles, v)
		}
	}

	return uniqueTitles
}
