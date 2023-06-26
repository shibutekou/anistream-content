package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vgekko/ani-go/config"
	v1 "github.com/vgekko/ani-go/internal/controller/http/v1"
	"github.com/vgekko/ani-go/internal/usecase"
	"github.com/vgekko/ani-go/internal/usecase/repo"
	"github.com/vgekko/ani-go/internal/usecase/webapi"
	"github.com/vgekko/ani-go/pkg/httpserver"
	"github.com/vgekko/ani-go/pkg/logger"
	"github.com/vgekko/ani-go/pkg/postgres"
	"github.com/vgekko/ani-go/pkg/redis"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(cfg *config.Config) {
	l := logger.New()

	// initialize postgres
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal("failed to connect to database: " + err.Error())
	}
	defer pg.Close()

	// initialize redis
	redisClient := redis.New()

	// initialize usecases
	infoUseCase := usecase.NewInfoUseCase(webapi.New(), repo.NewInfoRedis(redisClient, time.Hour*12))
	linkUseCase := usecase.NewLinkUseCase(webapi.New())

	// HTTP server
	engine := gin.New()
	v1.NewRouter(engine, infoUseCase, linkUseCase, l)

	httpServer := httpserver.New(engine)

	// waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app.Run: signal: " + s.String())
	case err := <-httpServer.Notify():
		l.Error(fmt.Errorf("app.Run: httpServer.Notify: %w", err).Error())
	}

	// shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app.Run: httpServer.Shutdown: %w", err).Error())
	}
}
