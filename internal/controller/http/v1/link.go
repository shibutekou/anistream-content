package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/vgekko/ani-go/internal/usecase"
	"github.com/vgekko/ani-go/pkg/apperror"
	"go.uber.org/zap"
	"net/http"
)

type linkRoutes struct {
	uc usecase.Link
	l  *zap.Logger
}

func newLinkRoutes(handler *gin.RouterGroup, uc usecase.Link, l *zap.Logger) {
	r := &linkRoutes{uc: uc, l: l}

	h := handler.Group("/link")
	{
		h.GET("kinopoisk", r.linkByKinopoiskID)
		h.GET("shikimori", r.linkByShikimoriID)
		h.GET("imdb", r.linkByIMDbID)
		h.GET("title", r.linkByTitleNameHandler)
	}
}

func (r *linkRoutes) linkByKinopoiskID(c *gin.Context) {
	id := c.Query("kinopoisk_id")
	if id == "" {
		r.l.Info("parameter kinopoisk_id is required")
		c.JSON(http.StatusBadRequest, "parameter kinopoisk_id is required")
	}

	link, err := r.uc.ByKinopoiskID(id)
	if err != nil {
		if errors.Is(err, apperror.ErrTitleNotFound) {
			r.l.Info("no such title by given kinopoisk id: " + err.Error())

			c.JSON(http.StatusNotFound, err.Error())
		}
	}

	c.JSON(http.StatusOK, link)
}

func (r *linkRoutes) linkByShikimoriID(c *gin.Context) {
	id := c.DefaultQuery("shikimori_id", "")
	if id == "" {
		r.l.Info("parameter shikimori_id is required")
		c.JSON(http.StatusBadRequest, "parameter shikimori_id is required")
	}

	link, err := r.uc.ByShikimoriID(id)
	if err != nil {
		if errors.Is(err, apperror.ErrTitleNotFound) {
			r.l.Info("no such title by given shikimori id: " + err.Error())

			c.JSON(http.StatusNotFound, err.Error())
		}
	}

	c.JSON(http.StatusOK, link)
}

func (r *linkRoutes) linkByIMDbID(c *gin.Context) {
	id := c.DefaultQuery("imdb_id", "")
	if id == "" {
		r.l.Info("parameter imdb_id is required")
		c.JSON(http.StatusBadRequest, "parameter imdb_id is required")
	}

	link, err := r.uc.ByIMDbID(id)
	if err != nil {
		if errors.Is(err, apperror.ErrTitleNotFound) {
			r.l.Info("no such title by given imdb id: " + err.Error())

			c.JSON(http.StatusNotFound, err.Error())
		}
	}

	c.JSON(http.StatusOK, link)
}

func (r *linkRoutes) linkByTitleNameHandler(c *gin.Context) {
	title := c.DefaultQuery("title", "")
	if title == "" {
		r.l.Info("parameter title is required")
		c.JSON(http.StatusBadRequest, "parameter title is required")
	}

	link, err := r.uc.ByTitleName(title)
	if err != nil {
		if errors.Is(err, apperror.ErrTitleNotFound) {
			r.l.Info("no such title by given title name: " + err.Error())

			c.JSON(http.StatusNotFound, err.Error())
		}
	}

	c.JSON(http.StatusOK, link)
}
