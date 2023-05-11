package delivery

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func (h *Handler) KinopoiskIDHandler(c *fiber.Ctx) error {
	kinopoiskID := c.Query("kinopoisk_id")

	link, err := h.kodik.ByKinopoiskID(kinopoiskID)
	if err != nil {
		log.Printf("error while getting link by kinopoisk id %s: %v", kinopoiskID, err)
		return err
	}

	if err := c.JSON(link); err != nil {
		return err
	}

	return nil
}
