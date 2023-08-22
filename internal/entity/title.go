package entity

type TitleContent struct {
	Link             string `json:"link"`
	Title            string `json:"title"`
	TitleOrig        string `json:"title_orig"`
	OtherTitle       string `json:"other_title"`
	Year             int32  `json:"year"`
	KinopoiskID      string `json:"kinopoisk_id"`
	ShikimoriID      string `json:"shikimori_id"`
	IMDbID           string `json:"imdb_id"`
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
	Screenshots   []string `json:"screenshots"`
}

type TitleFilter struct {
	Opt string
	Val string
}
