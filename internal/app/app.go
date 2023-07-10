package app

import (
	"github.com/gin-gonic/gin"
	"github.com/vgekko/ani-go/config"
	v1 "github.com/vgekko/ani-go/internal/controller/http/v1"
	"github.com/vgekko/ani-go/internal/usecase"
	"github.com/vgekko/ani-go/internal/usecase/repo"
	"github.com/vgekko/ani-go/internal/usecase/webapi"
	"github.com/vgekko/ani-go/pkg/httpserver"
	"github.com/vgekko/ani-go/pkg/logger/sl"
	"github.com/vgekko/ani-go/pkg/postgres"
	"github.com/vgekko/ani-go/pkg/redis"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	cfg := config.Load()

	// initialize slog logger
	log := sl.New(cfg.Env)

	// initialize postgres
	pg, err := postgres.New(cfg.Postgres.URL, postgres.MaxPoolSize(cfg.Postgres.PoolMax))
	if err != nil {
		log.Error("failed to init postgres", sl.Err(err))
	}
	defer pg.Close()

	// initialize redis
	redisClient := redis.New()

	// initialize usecases
	infoUseCase := usecase.NewInfoUseCase(webapi.New(), repo.NewInfoRedis(redisClient, cfg.Redis.InfoCacheTTL))
	linkUseCase := usecase.NewLinkUseCase(webapi.New())

	// HTTP server
	engine := gin.New()
	v1.NewRouter(engine, infoUseCase, linkUseCase, log)

	httpServer := httpserver.New(engine)

	// waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("app.Run: signal: " + s.String())
	case err := <-httpServer.Notify():
		log.Error("app.Run: notify: ", sl.Err(err))
	}

	// shutdown
	err = httpServer.Shutdown()
	if err != nil {
		log.Error("app.Run: shutdown: ", sl.Err(err))
	}
}
