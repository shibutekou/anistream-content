package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/vgekko/ani-go/internal/usecase"
	"go.uber.org/zap"
)

func NewRouter(handler *gin.Engine, i usecase.Info, l usecase.Link, log *zap.Logger) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	v1 := handler.Group("/v1")
	{
		newInfoRoutes(v1, i, log)
		newLinkRoutes(v1, l, log)
	}
}
