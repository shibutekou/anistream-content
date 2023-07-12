package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/vgekko/ani-go/internal/controller/http/v1/middleware"
	"golang.org/x/exp/slog"
)

type errorResponse struct {
	Error string `json:"error"`
}

type successResponse struct {
	Message string `json:"message"`
}

// TODO: подобрать красивое решение
func newErrorResponse(c *gin.Context, statusCode int, message string, log *slog.Logger) {
	log = log.With(
		slog.String("method", c.Request.Method),
		slog.Int("status", statusCode),
		slog.String("path", c.Request.URL.Path),
		slog.String("remote_addr", c.Request.RemoteAddr),
		slog.String("user_agent", c.Request.UserAgent()),
		slog.String("request_id", middleware.GetRequestID(c)),
	)

	log.Error(message)
	c.JSON(statusCode, errorResponse{message})
}

func newSuccessResponse(c *gin.Context, statusCode int, log *slog.Logger) {
	log = log.With(
		slog.String("method", c.Request.Method),
		slog.Int("status", statusCode),
		slog.String("path", c.Request.URL.Path),
		slog.String("remote_addr", c.Request.RemoteAddr),
		slog.String("user_agent", c.Request.UserAgent()),
		slog.String("request_id", middleware.GetRequestID(c)),
	)

	log.Info("ok")
}
