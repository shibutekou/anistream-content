package v1

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/vgekko/ani-go/internal/usecase"
	"golang.org/x/exp/slog"
)

func NewRouter(handler *gin.Engine, i usecase.Info, l usecase.Link, log *slog.Logger) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
	handler.Use(cors.New(cors.Config{
		AllowAllOrigins:        true,
		AllowMethods:           []string{"PUT", "PATCH", "POST", "DELETE", "GET"},
		AllowHeaders:           []string{"Origin", "Authorization", "Content-Type", "Accept-Encoding"},
		AllowCredentials:       true,
	}))

	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	v1 := handler.Group("/v1")
	{
		search := v1.Group("/search")
		{
			newInfoRoutes(search, i, log)
			newLinkRoutes(search, l, log)
		}

	}
}
