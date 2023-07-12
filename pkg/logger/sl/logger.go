package sl

import (
	"io"
	stdlog "log"
	"os"

	"golang.org/x/exp/slog"
)

const (
	EnvLocal = "local"
	EnvProd  = "prod"
)

func New(env string) *slog.Logger {
	var log *slog.Logger

	file, err := os.OpenFile(os.Getenv("ANIGO_LOGS"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		stdlog.Fatal(err)
	}

	switch env {
	case EnvLocal:
		log = setupPrettySlog(os.Stdout)
	case EnvProd:
		log = slog.New(
			slog.NewJSONHandler(io.MultiWriter(os.Stdout, file), &slog.HandlerOptions{Level: slog.LevelInfo}),
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
