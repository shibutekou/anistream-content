package normalize

import (
	"fmt"
	"github.com/vgekko/ani-go/internal/entity"
	"strings"

	"golang.org/x/exp/slices"
)

var validFilterParams = []string{"kinopoisk_id", "shikimori_id", "imdb_id", "worldart_id"}

func FilterParams(urlParams string) (entity.TitleFilter, error) {
	x := strings.Split(urlParams, "=")
	if !slices.Contains(validFilterParams, x[0]) {
		return entity.TitleFilter{}, fmt.Errorf("invalid search parameter")
	}

	return entity.TitleFilter{Option: x[0], Value:  x[1]}, nil
}

func URL(link string) string {
	return fmt.Sprintf("http://%s", strings.TrimLeft(link, "/"))
}
