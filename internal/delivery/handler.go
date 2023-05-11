package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vgekko/ani-go/internal/domain"
)

type Handler struct {
	kodik *domain.Kodik
}

func NewHandler(kodik *domain.Kodik) *Handler {
	return &Handler{kodik: kodik}
}

func (h *Handler) InitRoutes() *fiber.App {
	app := fiber.New()

	anime := app.Group("/anime")
	anime.Get("/link", h.KinopoiskIDHandler)
	//anime.Get("/search")

	return app
}
