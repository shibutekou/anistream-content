package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vgekko/ani-go/internal/domain"
	"go.uber.org/zap"
)

type Handler struct {
	link *domain.Link
	log  *zap.SugaredLogger
}

func NewHandler(link *domain.Link, log *zap.SugaredLogger) *Handler {
	return &Handler{link: link, log: log}
}

func (h *Handler) InitRoutes() *fiber.App {
	app := fiber.New()

	title := app.Group("/title")
	{
		title.Get("/link", h.LinkByIDHandler)
	}

	return app
}
