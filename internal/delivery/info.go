package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vgekko/ani-go/pkg/apperror"
	"net/http"
)

func (h *Handler) InfoByIDHandler(c *fiber.Ctx) error {
	service := h.dtm.determineService(c)
	if service == "" {
		h.log.Info("filter not passed or does not exist")

		c.Status(http.StatusBadRequest)
		return c.JSON("filter not passed or does not exist")
	}

	id := c.Query(service)

	titleInfos, err := h.service.Info.ByID(service, id)
	if err != nil {
		if err == apperror.ErrTitleNotFound {
			h.log.Infof("no such title by %s %s: %v", service, id, err)

			c.Status(http.StatusNotFound)
			return c.JSON(err.Error())
		}
	}

	return c.JSON(titleInfos)
}

func (h *Handler) InfoByTitleNameHandler(c *fiber.Ctx) error {
	title := c.Query("title")
	if title == "" {
		h.log.Info("parameter title is required")

		c.Status(http.StatusBadRequest)
		return c.JSON("parameter title is required")
	}

	titleInfos, err := h.service.Info.ByTitleName(title)
	if err != nil {
		if err == apperror.ErrTitleNotFound {
			h.log.Infof("no such title by name %s: %v", title, err)

			c.Status(http.StatusNotFound)
			return c.JSON(err.Error())
		}
	}

	return c.JSON(titleInfos)
}
