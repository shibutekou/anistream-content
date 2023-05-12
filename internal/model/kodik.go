package model

import "time"

type KodikAPI struct {
	Time  string `json:"time,omitempty"`
	Total int    `json:"total,omitempty"`
	Results
}

type Results []struct {
	ID            string `json:"id,omitempty"`
	Link          string `json:"link,omitempty"`
	Title         string `json:"title,omitempty"`
	TitleOrig     string `json:"title_orig,omitempty"`
	OtherTitle    string `json:"other_title"`
	Translation   `json:"translation"`
	Year          int       `json:"year,omitempty"`
	LastSeason    int       `json:"last_season,omitempty"`
	LastEpisode   int       `json:"last_episode,omitempty"`
	EpisodesCount int       `json:"episodes_count,omitempty"`
	KinopoiskID   string    `json:"kinopoisk_id,omitempty"`
	ImdbID        string    `json:"imdb_id,omitempty"`
	ShikimoriID   string    `json:"shikimori_id,omitempty"`
	Quality       string    `json:"quality,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Seasons       `json:"seasons"`
	Screenshots   []string `json:"screenshots,omitempty"`
}

type Translation struct {
	ID    int    `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Voice string `json:"voice,omitempty"`
}

type Seasons map[string]struct {
	Link     string `json:"link"`
	Episodes `json:"episodes"`
}

type Episodes map[string]string
