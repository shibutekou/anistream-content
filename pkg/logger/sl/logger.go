package sl

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
		log = setupPrettySlog(os.Stdout)
	case EnvProd:
		log = slog.New(
			slog.NewJSONHandler(io.Writer(os.Stdout), &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
