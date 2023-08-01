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

type linkRoutes struct {
	uc  usecase.LinkUseCase
	log *slog.Logger
}

func newLinkRoutes(handler *gin.RouterGroup, uc usecase.LinkUseCase, log *slog.Logger) {
	r := &linkRoutes{uc: uc, log: log}

	handler.GET("/link", r.search)
}

func (r *linkRoutes) search(c *gin.Context) {
	params, err := url.ParseQuery(c.Request.URL.RawQuery)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error(), r.log)
		return
	}

	filter, err := normalize.FilterParams(params.Encode())
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error(), r.log)
		return
	}

	link, err := r.uc.Search(filter)
	if err != nil {
		if errors.Is(err, apperror.ErrTitleNotFound) {
			newErrorResponse(c, http.StatusNotFound, err.Error(), r.log)
			return
		}

		newErrorResponse(c, http.StatusInternalServerError, err.Error(), r.log)
		return
	}

	c.JSON(http.StatusOK, link)
}
