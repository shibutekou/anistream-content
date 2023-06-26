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

type infoRoutes struct {
	uc usecase.Info
	l  *zap.Logger
}

func newInfoRoutes(handler *gin.RouterGroup, uc usecase.Info, l *zap.Logger) {
	r := &infoRoutes{uc: uc, l: l}

	h := handler.Group("/info")
	{
		h.GET("search", r.search)
	}
}

func (r *infoRoutes) search(c *gin.Context) {
	params, err := url.ParseQuery(c.Request.URL.RawQuery)
	if err != nil {
		r.l.Info("infoRoutes.Search: " + err.Error())
		c.JSON(http.StatusInternalServerError, "something went wrong")
	}

	option, value := normalize.Params(params.Encode(), "=")

	titleInfos, err := r.uc.Search(option, value)
	if err != nil {
		if errors.Is(err, apperror.ErrTitleNotFound) {
			r.l.Info("infoRoutes.Search: " + err.Error())
			c.JSON(http.StatusNotFound, "title with given parameters not found")

			return
		}
	}

	c.JSON(http.StatusOK, titleInfos)
}
