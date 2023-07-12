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

type infoRoutes struct {
	uc  usecase.Info
	log *slog.Logger
}

func newInfoRoutes(handler *gin.RouterGroup, uc usecase.Info, log *slog.Logger) {
	r := &infoRoutes{uc: uc, log: log}

	handler.GET("/info", r.search)
}

func (r *infoRoutes) search(c *gin.Context) {
	params, err := url.ParseQuery(c.Request.URL.RawQuery)
	if err != nil {
		r.log.Error("infoRoutes.Search: ", sl.Err(err))
		c.JSON(http.StatusInternalServerError, "something went wrong")
		return
	}

	option, value, err := normalize.Params(params.Encode())
	if err != nil {
		r.log.Error("normalize.Params: ", sl.Err(err))
		c.JSON(http.StatusBadRequest, "invalid search parameter")
		return
	}

	titleInfos, err := r.uc.Search(option, value)
	if err != nil {
		if errors.Is(err, apperror.ErrTitleNotFound) {
			//r.log.Warn("infoRoutes.Search: ", sl.Err(err))
			//c.JSON(http.StatusNotFound, "title with given parameters not found")
			errorResponse(c, "title with given parameters not found", http.StatusNotFound, r.log)
			return
		}

		r.log.Error("infoRoute.search: ", sl.Err(err))
		c.JSON(http.StatusInternalServerError, "something went wrong")
		return
	}

	c.JSON(http.StatusOK, titleInfos)
}
