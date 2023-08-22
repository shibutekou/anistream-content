package entity

type KodikAPI struct {
	Time    string `json:"time,omitempty"`
	Total   int    `json:"total,omitempty"`
	Results `json:"results"`
}

type Results []struct {
	Link         string `json:"link"`
	Title        string `json:"title"`
	TitleOrig    string `json:"title_orig"`
	OtherTitle   string `json:"other_title"`
	Year         int32  `json:"year"`
	KinopoiskID  string `json:"kinopoisk_id"`
	ShikimoriID  string `json:"shikimori_id"`
	IMDbID       string `json:"imdb_id"`
	MaterialData `json:"material_data"`
	Screenshots  []string `json:"screenshots"`
}

type MaterialData struct {
	AnimeStatus      string `json:"anime_status"`
	AnimeDescription string `json:"anime_description"`
	PosterURL        string `json:"poster_url"`
	Duration         int32  `json:"duration"`

	KinopoiskRating float64 `json:"kinopoisk_rating"`
	IMDbRating      float64 `json:"imdb_rating"`
	ShikimoriRating float64 `json:"shikimori_rating"`

	PremiereWorld string   `json:"premiere_world"`
	EpisodesTotal int32    `json:"episodes_total"`
	Writers       []string `json:"writers"`
}
