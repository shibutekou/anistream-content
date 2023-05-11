package delivery

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func (h *Handler) ByIDHandler(c *fiber.Ctx) error {
	service := determineService(c)
	id := c.Query(service)

	link, err := h.kodik.ByServiceID(service, id)
	if err != nil {
		log.Printf("error while getting link by %s %s: %v", service, id, err)
		return err
	}

	if err := c.JSON(link); err != nil {
		return err
	}

	return nil
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
