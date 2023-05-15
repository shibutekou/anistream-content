package model

type TitleInfo struct {
	Title         string   `json:"title,omitempty"`
	TitleOrig     string   `json:"title_orig,omitempty"`
	OtherTitle    string   `json:"other_title"`
	Year          int      `json:"year,omitempty"`
	LastSeason    int      `json:"last_season,omitempty"`
	LastEpisode   int      `json:"last_episode,omitempty"`
	EpisodesCount int      `json:"episodes_count,omitempty"`
	KinopoiskID   string   `json:"kinopoisk_id,omitempty"`
	ImdbID        string   `json:"imdb_id,omitempty"`
	ShikimoriID   string   `json:"shikimori_id,omitempty"`
	Screenshots   []string `json:"screenshots"`
}
