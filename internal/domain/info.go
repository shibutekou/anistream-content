package domain

import (
	"encoding/json"
	"fmt"
	"github.com/vgekko/ani-go/internal/model"
	"github.com/vgekko/ani-go/pkg/apperror"
	"net/http"
)

type InfoServiceImpl struct {
	token      string
	client     http.Client
	conversion conversionService
}

func NewInfoService(token string, client http.Client, conversion conversionService) *InfoServiceImpl {
	return &InfoServiceImpl{
		token:      token,
		client:     client,
		conversion: conversion,
	}
}

func (i *InfoServiceImpl) ByID(service, id string) ([]model.TitleInfo, error) {
	var kodikResponse model.KodikAPI
	var titleInfos []model.TitleInfo

	url := fmt.Sprintf("https://kodikapi.com/search?token=%s&%s=%s", i.token, service, id)

	resp, err := i.client.Get(url)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&kodikResponse); err != nil {
		return nil, err
	}

	if len(kodikResponse.Results) == 0 {
		return nil, apperror.ErrTitleNotFound
	}

	titleInfos = i.conversion.toTitleInfo(&kodikResponse)

	return titleInfos, nil
}

func (i *InfoServiceImpl) ByTitleName(title string) ([]model.TitleInfo, error) {
	var kodikResponse model.KodikAPI
	var titleInfos []model.TitleInfo

	url := fmt.Sprintf("https://kodikapi.com/search?token=eb045488556bdd7adae2a023a8c786b8&title=%s", title)

	resp, err := i.client.Get(url)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(resp.Body).Decode(&kodikResponse); err != nil {
		return nil, err
	}

	if len(kodikResponse.Results) == 0 {
		return nil, apperror.ErrTitleNotFound
	}

	titleInfos = i.conversion.toTitleInfo(&kodikResponse)

	return titleInfos, nil
}
