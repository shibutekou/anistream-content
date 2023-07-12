package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func Slogger(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		log = log.With(
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.URL.Path),
			slog.String("remote_addr", c.Request.RemoteAddr),
			slog.String("user_agent", c.Request.UserAgent()),
			slog.String("request_id", GetRequestID(c)),
		)

		log.Debug("request completed")
	}
}
