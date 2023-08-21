package v1

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/vgekko/anistream-content/internal/controller/http/v1/middleware"
	"github.com/vgekko/anistream-content/internal/usecase"
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

	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	v1 := handler.Group("/v1")
	{
		search := v1.Group("/search")
		{
			newInfoRoutes(search, uc.InfoUseCase, log)
			newLinkRoutes(search, uc.LinkUseCase, log)
		}

	}
}
