package v1

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/vgekko/ani-go/internal/controller/http/v1/middleware"
	"github.com/vgekko/ani-go/internal/usecase"
	"golang.org/x/exp/slog"
)

func NewRouter(handler *gin.Engine, uc *usecase.UseCase, log *slog.Logger) {
	handler.Use(gin.Recovery())
	handler.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT", "PATCH", "POST", "DELETE", "GET"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Accept-Encoding"},
		AllowCredentials: true,
	}))

	handler.Use(middleware.RequestID())
	handler.Use(middleware.Slogger(log))

	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	v1 := handler.Group("/v1")
	{
		newAuthRoutes(v1, uc.Auth, log)

		search := v1.Group("/search")
		{
			newInfoRoutes(search, uc.Info, log)
			newLinkRoutes(search, uc.Link, log)
		}

	}
}
