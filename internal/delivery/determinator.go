package delivery

import "github.com/gofiber/fiber/v2"

type DeterminatorImpl struct{}

func NewDeterminator() *DeterminatorImpl {
	return &DeterminatorImpl{}
}

type determinator interface {
	determineService(c *fiber.Ctx) string
}

func (d *DeterminatorImpl) determineService(c *fiber.Ctx) string {
	availableServices := []string{"kinopoisk_id", "imdb_id", "shikimori_id", "mdl_id"}

	for _, s := range availableServices {
		if c.Query(s) != "" {
			return s
		}
	}

	return ""
}
