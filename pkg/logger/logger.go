package logger

import "go.uber.org/zap"

func Sugared() *zap.SugaredLogger {
	l, _ := zap.NewProduction()
	defer l.Sync()

	sugar := l.Sugar()

	return sugar
}
