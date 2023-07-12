package v1

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/vgekko/ani-go/internal/usecase"
	"github.com/vgekko/ani-go/pkg/apperror"
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
		newErrorResponse(c, http.StatusBadRequest, err.Error(), r.log)
		return
	}

	option, value, err := normalize.Params(params.Encode())
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error(), r.log)
		return
	}

	titleInfos, err := r.uc.Search(option, value)
	if err != nil {
		if errors.Is(err, apperror.ErrTitleNotFound) {
			newErrorResponse(c, http.StatusNotFound, err.Error(), r.log)
			return
		}

		newErrorResponse(c, http.StatusInternalServerError, err.Error(), r.log)
		return
	}

	newSuccessResponse(c, http.StatusOK, r.log)
	c.JSON(http.StatusOK, titleInfos)
}
