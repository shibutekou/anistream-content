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

	stdlog "log"
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
		stdlog.Fatalf("app.Run: %s", err)
	}

	log.Info("starting grpc server")
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			stdlog.Fatalf("app.Run: %s", err.Error())
		}
	}()

	// waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	<-interrupt

	// stopping gRPC server
	grpcServer.GracefulStop()

	// closing listener
	err = lis.Close()
	if err != nil {
		stdlog.Fatalf("app.Run: lis.Close: %v", err.Error())
	}
}
