package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vgekko/ani-go/internal/domain"
	"go.uber.org/zap"
)

type Handler struct {
	service *domain.Service
	dtm     determinator
	log     *zap.SugaredLogger
}

func NewHandler(service *domain.Service, log *zap.SugaredLogger) *Handler {
	dtm := NewDeterminator()
	return &Handler{service: service, dtm: dtm, log: log}
}

func (h *Handler) InitRoutes() *fiber.App {
	app := fiber.New()

	title := app.Group("/title")
	{
		title.Get("/link/id", h.LinkByIDHandler)
		title.Get("/link/name", h.LinkByTitleNameHandler)
		title.Get("/info/id", h.InfoByIDHandler)
		title.Get("/info/name", h.InfoByTitleNameHandler)
	}

	return app
}
