package sl

import (
	"golang.org/x/exp/slog"
	"os"
)

const (
	EnvLocal = "local"
	EnvProd = "prod"
)

func New(env string) *slog.Logger {
	 var log *slog.Logger

	 switch env {
	 case EnvLocal:
		 log = slog.New(
			 slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		 )
	 case EnvProd:
		 log = slog.New(
			 slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
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