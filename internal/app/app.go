package app

import (
	"github.com/gin-gonic/gin"
	"github.com/vgekko/anistream-content/config"
	controllerGrpc "github.com/vgekko/anistream-content/internal/controller/grpc"
	v1 "github.com/vgekko/anistream-content/internal/controller/http/v1"
	"github.com/vgekko/anistream-content/internal/repository"
	"github.com/vgekko/anistream-content/internal/usecase"
	"github.com/vgekko/anistream-content/internal/webapi"
	"github.com/vgekko/anistream-content/pkg/cache"
	"github.com/vgekko/anistream-content/pkg/httpserver"
	"github.com/vgekko/anistream-content/pkg/logger/sl"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	cfg := config.Load()

	// initialize slog logger
	log := sl.New(cfg.Env)

	// initialize cache
	cacher := cache.New(cfg.Cache)

	// web api
	webAPI := webapi.NewWebAPI()

	// repositories
	repo := repository.NewRepository(cacher)

	// use cases
	useCase := usecase.NewUseCase(repo, webAPI)

	// HTTP server
	engine := gin.New()
	v1.NewRouter(engine, useCase, log)

	// starting HTTP server
	log.Info("starting http server")
	httpServer := httpserver.Start(engine)

	// starting gRPC server
	grpcServer := grpc.NewServer()
	controllerGrpc.NewContentServerGrpc(grpcServer, useCase.InfoUseCase, log)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Error("could not listen tcp :50051: ", sl.Err(err))
	}

	log.Info("starting grpc server")
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Error("app.Run: grpc: ", sl.Err(err))
		}
	}()

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
