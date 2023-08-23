package logger

import (
	"io"
	"os"

	"log/slog"
)

const (
	EnvLocal = "local"
	EnvProd  = "prod"
)

func New(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case EnvLocal:
		log = slog.New(slog.NewTextHandler(io.Writer(os.Stdout), &slog.HandlerOptions{Level: slog.LevelDebug}))
	case EnvProd:
		log = slog.New(
			slog.NewJSONHandler(io.Writer(os.Stdout), &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
