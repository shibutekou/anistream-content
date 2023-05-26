package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/vgekko/ani-go/internal/usecase"
	"github.com/vgekko/ani-go/pkg/apperror"
	"go.uber.org/zap"
	"net/http"
)

type infoRoutes struct {
	uc usecase.Info
	l  *zap.Logger
}

func newInfoRoutes(handler *gin.RouterGroup, uc usecase.Info, l *zap.Logger) {
	r := &infoRoutes{uc: uc, l: l}

	h := handler.Group("/info")
	{
		h.GET("kinopoisk", r.infoByKinopoiskID)
		h.GET("shikimori", r.infoByShikimoriID)
		h.GET("imdb", r.infoByIMDbID)
		h.GET("title", r.InfoByTitleNameHandler)
	}
}

func (r *infoRoutes) infoByKinopoiskID(c *gin.Context) {
	id := c.Query("kinopoisk_id")
	if id == "" {
		r.l.Info("parameter kinopoisk_id is required")
		c.JSON(http.StatusBadRequest, "parameter kinopoisk_id is required")
	}

	titleInfos, err := r.uc.ByKinopoiskID(id)
	if err != nil {
		if err == apperror.ErrTitleNotFound {
			r.l.Info("no such title by given kinopoisk id: " + err.Error())

			c.JSON(http.StatusNotFound, err.Error())
		}
	}

	c.JSON(http.StatusOK, titleInfos)
}

func (r *infoRoutes) infoByShikimoriID(c *gin.Context) {
	id := c.Query("shikimori_id")
	if id == "" {
		r.l.Info("parameter shikimori_id is required")
		c.JSON(http.StatusBadRequest, "parameter shikimori_id is required")
	}

	titleInfos, err := r.uc.ByShikimoriID(id)
	if err != nil {
		if err == apperror.ErrTitleNotFound {
			r.l.Info("no such title by given shikimori id: " + err.Error())

			c.JSON(http.StatusNotFound, err.Error())
		}
	}

	c.JSON(http.StatusOK, titleInfos)
}

func (r *infoRoutes) infoByIMDbID(c *gin.Context) {
	id := c.Query("imdb_id")
	if id == "" {
		r.l.Info("parameter imdb_id is required")
		c.JSON(http.StatusBadRequest, "parameter imdb_id is required")
	}

	titleInfos, err := r.uc.ByIMDbID(id)
	if err != nil {
		if err == apperror.ErrTitleNotFound {
			r.l.Info("no such title by given imdb id: " + err.Error())

			c.JSON(http.StatusNotFound, err.Error())
		}
	}

	c.JSON(http.StatusOK, titleInfos)
}

func (r *infoRoutes) InfoByTitleNameHandler(c *gin.Context) {
	title := c.Query("title")
	if title == "" {
		r.l.Info("parameter title is required")
		c.JSON(http.StatusBadRequest, "parameter title is required")
	}

	titleInfos, err := r.uc.ByTitleName(title)
	if err != nil {
		if err == apperror.ErrTitleNotFound {
			r.l.Info("no such title by given title name: " + err.Error())

			c.JSON(http.StatusNotFound, err.Error())
		}
	}

	c.JSON(http.StatusOK, titleInfos)
}
