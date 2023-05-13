package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vgekko/ani-go/internal/domain"
	"go.uber.org/zap"
)

type Handler struct {
	service *domain.Service
	log     *zap.SugaredLogger
}

func NewHandler(service *domain.Service, log *zap.SugaredLogger) *Handler {
	return &Handler{service: service, log: log}
}

func (h *Handler) InitRoutes() *fiber.App {
	app := fiber.New()

	title := app.Group("/title")
	{
		title.Get("/link/id", h.LinkByIDHandler)
		title.Get("/link/title", h.LinkByTitleNameHandler)
		title.Get("/info/id", h.InfoByIDHandler)
		title.Get("/info/title", h.InfoByTitleNameHandler)
	}

	return app
}

func determineService(c *fiber.Ctx) string {
	availableServices := []string{"kinopoisk_id", "imdb_id", "shikimori_id", "mdl_id"}

	for _, s := range availableServices {
		if c.Query(s) != "" {
			return s
		}
	}

	return ""
}
