// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/vgekko/ani-go/internal/entity"
)

type LinkPayload struct {
	Link entity.Link `json:"link"`
}

type TitleFilter struct {
	Option string `json:"option"`
	ID     string `json:"id"`
}

type TitleInfoPayload struct {
	TitleInfo []entity.TitleInfo `json:"titleInfo"`
}