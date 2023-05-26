package logger

import "go.uber.org/zap"

func New() *zap.Logger {
	l, _ := zap.NewProduction()
	defer l.Sync()

	return l
}
