package entity

import (
	"encoding/json"
	"github.com/bytedance/sonic/encoder"
)

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

type KodikAPI struct {
	Time    string `json:"time,omitempty"`
	Total   int    `json:"total,omitempty"`
	Results `json:"results"`
}

type Results []struct {
	Link string `json:"link"`
	Title       string   `json:"title,omitempty"`
	TitleOrig   string   `json:"title_orig,omitempty"`
	OtherTitle  string   `json:"other_title"`
	Year        int      `json:"year,omitempty"`
	KinopoiskID string   `json:"kinopoisk_id,omitempty"`
	ShikimoriID string   `json:"shikimori_id,omitempty"`
	IMDbID      string   `json:"imdb_id,omitempty"`
	Screenshots []string `json:"screenshots"`
}

type TitleInfos []TitleInfo

func (tis TitleInfos) MarshalBinary() (data []byte, err error) {
	b, err := encoder.Encode(tis, encoder.CompactMarshaler)

	return b, nil
}

func (tis TitleInfos) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &tis)
}
