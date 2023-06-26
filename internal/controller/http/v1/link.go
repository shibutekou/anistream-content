package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/vgekko/ani-go/internal/usecase"
	"github.com/vgekko/ani-go/pkg/apperror"
	"github.com/vgekko/ani-go/pkg/normalize"
	"go.uber.org/zap"
	"net/http"
	"net/url"
)

type linkRoutes struct {
	uc usecase.Link
	l  *zap.Logger
}

func newLinkRoutes(handler *gin.RouterGroup, uc usecase.Link, l *zap.Logger) {
	r := &linkRoutes{uc: uc, l: l}

	h := handler.Group("/link")
	{
		h.GET("search", r.search)
	}
}

func (r *linkRoutes) search(c *gin.Context) {
	params, err := url.ParseQuery(c.Request.URL.RawQuery)
	if err != nil {
		r.l.Info("linkRoutes.search: " + err.Error())
		c.JSON(http.StatusInternalServerError, "something went wrong")
	}

	option, value := normalize.Params(params.Encode(), "=")

	link, err := r.uc.Search(option, value)
	if err != nil {
		if errors.Is(err, apperror.ErrTitleNotFound) {
			r.l.Info("linkRoutes.search: " + err.Error())
			c.JSON(http.StatusNotFound, "title with given parameters not found")

			return
		}
	}

	c.JSON(http.StatusOK, link)
}


