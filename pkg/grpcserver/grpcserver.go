package grpcserver

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func New(log *slog.Logger) *grpc.Server {
	logOpts := []logging.Option{logging.WithLogOnEvents(logging.StartCall, logging.FinishCall)}
	recOpts := []recovery.Option{recovery.WithRecoveryHandler(InterceptorRecovery())}

	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		logging.UnaryServerInterceptor(InterceptorLogging(log), logOpts...),
		recovery.UnaryServerInterceptor(recOpts...),
	))

	return grpcServer
}

func InterceptorLogging(log *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, level logging.Level, msg string, fields ...any) {
		log.Log(ctx, slog.LevelDebug, msg, fields...)
	})
}

func InterceptorRecovery() recovery.RecoveryHandlerFunc {
	return func(p any) (err error) {
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	}
}
