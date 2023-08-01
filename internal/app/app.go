package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"github.com/vgekko/ani-go/config"
	v1 "github.com/vgekko/ani-go/internal/controller/http/v1"
	redisRepository "github.com/vgekko/ani-go/internal/repository/redis"
	"github.com/vgekko/ani-go/internal/usecase"
	"github.com/vgekko/ani-go/internal/webapi"
	"github.com/vgekko/ani-go/pkg/httpserver"
	"github.com/vgekko/ani-go/pkg/logger/sl"
	redisClient "github.com/vgekko/ani-go/pkg/redis"
)

func Run() {
	cfg := config.Load()

	// initialize slog logger
	log := sl.New(cfg.Env)

	// initialize redis
	redis := redisClient.NewClient(cfg.Redis)
	defer redis.Close()

	// web api
	webAPI := webapi.NewWebAPI()

	// repositories
	redisRepo := redisRepository.NewRepositoryRedis(redis, cfg.Redis)

	// initialize usecases
	useCase := usecase.NewUseCase(redisRepo, webAPI)

	// HTTP server
	engine := gin.New()
	v1.NewRouter(engine, useCase, log)

	log.Info("starting http server")
	httpServer := httpserver.Start(engine)

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
	err := httpServer.Shutdown()
	if err != nil {
		log.Error("app.Run: shutdown: ", sl.Err(err))
	}
}
