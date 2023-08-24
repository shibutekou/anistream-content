package app

import (
	"fmt"
	"github.com/vgekko/anistream-content/config"
	controllerGrpc "github.com/vgekko/anistream-content/internal/controller/grpc"
	"github.com/vgekko/anistream-content/internal/repository"
	"github.com/vgekko/anistream-content/internal/usecase"
	"github.com/vgekko/anistream-content/internal/webapi"
	"github.com/vgekko/anistream-content/pkg/cache"
	"github.com/vgekko/anistream-content/pkg/grpcserver"
	"github.com/vgekko/anistream-content/pkg/logger"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	// initialize slog logger
	log := logger.New(cfg.Env)

	// initialize bigcache
	bc := cache.New(cfg.Cache)

	// web api
	webAPI := webapi.NewWebAPI()

	// repositories
	repo := repository.NewRepository(bc)

	// use cases
	useCase := usecase.NewUseCase(repo, webAPI)

	// starting gRPC server
	grpcServer := grpcserver.New(log)
	controllerGrpc.NewContentServerGrpc(grpcServer, useCase.ContentUseCase, log)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GRPC.Port))
	if err != nil {
		log.Error("could not listen tcp: ", err.Error())
	}

	log.Info("starting grpc server")
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Error("app.Run: grpc: ", err.Error())
		}
	}()

	// waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("app.Run: signal: " + s.String())
	}

	// stopping gRPC server
	grpcServer.GracefulStop()

	// closing listener
	err = lis.Close()
	if err != nil {
		log.Error("app.Run: lis.Close: ", err.Error())
	}
}
