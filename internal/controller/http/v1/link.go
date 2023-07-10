package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/vgekko/ani-go/internal/usecase"
	"github.com/vgekko/ani-go/pkg/apperror"
	"github.com/vgekko/ani-go/pkg/normalize"
	"golang.org/x/exp/slog"
	"net/http"
	"net/url"
)

type linkRoutes struct {
	uc usecase.Link
	log  *slog.Logger
}

func newLinkRoutes(handler *gin.RouterGroup, uc usecase.Link, log *slog.Logger) {
	r := &linkRoutes{uc: uc, log: log}

	handler.GET("/link", r.search)
}

func (r *linkRoutes) search(c *gin.Context) {
	params, err := url.ParseQuery(c.Request.URL.RawQuery)
	if err != nil {
		r.log.Info("linkRoutes.search: " + err.Error())
		c.JSON(http.StatusInternalServerError, "something went wrong")
	}

	option, value := normalize.Params(params.Encode(), "=")

	link, err := r.uc.Search(option, value)
	if err != nil {
		if errors.Is(err, apperror.ErrTitleNotFound) {
			r.log.Info("linkRoutes.search: " + err.Error())
			c.JSON(http.StatusNotFound, "title with given parameters not found")

			return
		}
	}

	c.JSON(http.StatusOK, link)
}


