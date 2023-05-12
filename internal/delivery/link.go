package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vgekko/ani-go/pkg/apperror"
	"net/http"
)

func (h *Handler) LinkByIDHandler(c *fiber.Ctx) error {
	service := determineService(c)
	if service == "" {
		h.log.Info("no such service")

		c.Status(http.StatusBadRequest)
		return c.JSON("no such service for searching title")
	}

	id := c.Query(service)

	link, err := h.link.ByID(service, id)
	if err != nil {
		if err == apperror.ErrTitleNotFound {
			h.log.Infof("no such title by %s %s: %v", service, id, err)

			c.Status(http.StatusInternalServerError)
			return c.JSON(err.Error())
		}
	}

	return c.JSON(link)
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
