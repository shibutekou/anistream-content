package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/vgekko/ani-go/internal/usecase"
	"go.uber.org/zap"
)

func NewRouter(handler *gin.Engine, i usecase.Info, l usecase.Link, log *zap.Logger) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	h := handler.Group("/v1")
	{
		newInfoRoutes(h, i, log)
		newLinkRoutes(h, l, log)
	}
}
