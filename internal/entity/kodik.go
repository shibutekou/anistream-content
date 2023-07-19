package entity

type TitleInfo struct {
	Title       string   `json:"title,omitempty"`
	TitleOrig   string   `json:"title_orig,omitempty"`
	OtherTitle  string   `json:"other_title,omitempty"`
	Year        int      `json:"year,omitempty"`
	KinopoiskID string   `json:"kinopoisk_id,omitempty"`
	ShikimoriID string   `json:"shikimori_id,omitempty"`
	IMDbID      string   `json:"imdb_id,omitempty"`
	Screenshots []string `json:"screenshots,omitempty"`
}

type Link struct {
	URL string `json:"url"`
}

type TitleFilter struct {
	Option string
	Value string
}

type KodikAPI struct {
	Time    string `json:"time,omitempty"`
	Total   int    `json:"total,omitempty"`
	Results `json:"results"`
}

type Results []struct {
	Link        string   `json:"link"`
	Title       string   `json:"title,omitempty"`
	TitleOrig   string   `json:"title_orig,omitempty"`
	OtherTitle  string   `json:"other_title"`
	Year        int      `json:"year,omitempty"`
	KinopoiskID string   `json:"kinopoisk_id,omitempty"`
	ShikimoriID string   `json:"shikimori_id,omitempty"`
	IMDbID      string   `json:"imdb_id,omitempty"`
	Screenshots []string `json:"screenshots"`
}
