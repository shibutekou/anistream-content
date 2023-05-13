package domain

import (
	"github.com/vgekko/ani-go/internal/model"
	"net/http"
)

type InfoServiceImpl struct {
	token  string
	client http.Client
}

func NewInfoService(token string, client http.Client) *InfoServiceImpl {
	return &InfoServiceImpl{
		token:  token,
		client: client,
	}
}

func (i *InfoServiceImpl) ByID(service, id string) ([]model.TitleInfo, error) {
	panic("implement")
}

func (i *InfoServiceImpl) ByTitleName(title string) ([]model.TitleInfo, error) {
	panic("implement")
}
