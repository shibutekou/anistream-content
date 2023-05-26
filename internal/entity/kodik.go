package entity

import "time"

type TitleInfo struct {
	Title       string   `json:"title,omitempty"`
	TitleOrig   string   `json:"title_orig,omitempty"`
	OtherTitle  string   `json:"other_title"`
	Year        int      `json:"year,omitempty"`
	KinopoiskID string   `json:"kinopoisk_id,omitempty"`
	ShikimoriID string   `json:"shikimori_id,omitempty"`
	IMDbID      string   `json:"imdb_id,omitempty"`
	Screenshots []string `json:"screenshots"`
}

type KodikAPI struct {
	Time    string `json:"time,omitempty"`
	Total   int    `json:"total,omitempty"`
	Results []struct {
		ID          string `json:"id,omitempty"`
		Link        string `json:"link,omitempty"`
		Title       string `json:"title,omitempty"`
		TitleOrig   string `json:"title_orig,omitempty"`
		OtherTitle  string `json:"other_title"`
		Translation struct {
			ID    int    `json:"id,omitempty"`
			Title string `json:"title,omitempty"`
			Voice string `json:"voice,omitempty"`
		} `json:"translation"`
		Year          int       `json:"year,omitempty"`
		LastSeason    int       `json:"last_season,omitempty"`
		LastEpisode   int       `json:"last_episode,omitempty"`
		EpisodesCount int       `json:"episodes_count,omitempty"`
		KinopoiskID   string    `json:"kinopoisk_id,omitempty"`
		IMDbID        string    `json:"imdb_id,omitempty"`
		ShikimoriID   string    `json:"shikimori_id,omitempty"`
		Quality       string    `json:"quality,omitempty"`
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"updated_at"`
		Seasons       map[string]struct {
			Link     string            `json:"link"`
			Episodes map[string]string `json:"episodes"`
		} `json:"seasons"`
		Screenshots []string `json:"screenshots,omitempty"`
	}
}

func (k *KodikAPI) ToTitleInfo() []TitleInfo {
	var ti TitleInfo
	titleInfos := make([]TitleInfo, len(k.Results), cap(k.Results))

	for _, v := range k.Results {
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

	return titleInfos
}
