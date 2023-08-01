package v1

import (
	graphqlhandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/vgekko/ani-go/graph"
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

	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	v1 := handler.Group("/v1")
	{
		v1.POST("/graphql", graphqlHandler(uc, log))

		search := v1.Group("/search")
		{
			newInfoRoutes(search, uc.InfoUseCase, log)
			newLinkRoutes(search, uc.LinkUseCase, log)
		}

	}
}

func graphqlHandler(uc *usecase.UseCase, log *slog.Logger) gin.HandlerFunc {
	h := graphqlhandler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Uc: uc, Log: log}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
