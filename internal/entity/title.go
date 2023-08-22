package entity

type TitleInfo struct {
	Link        string   `json:"link"`
	Title       string   `json:"title,omitempty"`
	TitleOrig   string   `json:"title_orig,omitempty"`
	OtherTitle  string   `json:"other_title,omitempty"`
	Year        int32    `json:"year,omitempty"`
	KinopoiskID string   `json:"kinopoisk_id,omitempty"`
	ShikimoriID string   `json:"shikimori_id,omitempty"`
	IMDbID      string   `json:"imdb_id,omitempty"`
	Screenshots []string `json:"screenshots,omitempty"`
}

type TitleFilter struct {
	Opt string
	Val string
}
