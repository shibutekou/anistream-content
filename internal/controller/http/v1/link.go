package v1

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/vgekko/ani-go/internal/usecase"
	"github.com/vgekko/ani-go/pkg/apperror"
	"github.com/vgekko/ani-go/pkg/logger/sl"
	"github.com/vgekko/ani-go/pkg/normalize"
	"golang.org/x/exp/slog"
)

type linkRoutes struct {
	uc  usecase.Link
	log *slog.Logger
}

func newLinkRoutes(handler *gin.RouterGroup, uc usecase.Link, log *slog.Logger) {
	r := &linkRoutes{uc: uc, log: log}

	handler.GET("/link", r.search)
}

func (r *linkRoutes) search(c *gin.Context) {
	params, err := url.ParseQuery(c.Request.URL.RawQuery)
	if err != nil {
		r.log.Error("linkRoutes.search: ", sl.Err(err))
		c.JSON(http.StatusInternalServerError, "something went wrong")
		return
	}

	option, value, err := normalize.Params(params.Encode())
	if err != nil {
		r.log.Error("normalize.Params: ", sl.Err(err))
		c.JSON(http.StatusBadRequest, "invalid search parameter")
		return
	}

	link, err := r.uc.Search(option, value)
	if err != nil {
		if errors.Is(err, apperror.ErrTitleNotFound) {
			r.log.Warn("linkRoutes.search: ", sl.Err(err))
			c.JSON(http.StatusNotFound, "title with given parameters not found")
			return
		}

		r.log.Error("linkRoute.search: ", sl.Err(err))
		c.JSON(http.StatusInternalServerError, "something went wrong")
		return
	}

	c.JSON(http.StatusOK, link)
}
