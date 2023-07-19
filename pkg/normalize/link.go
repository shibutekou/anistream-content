package normalize

import (
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

var validParams = []string{"kinopoisk_id", "shikimori_id", "imdb_id", "worldart_id"}

func Params(l string) (string, string, error) {
	x := strings.Split(l, "=")
	if !slices.Contains(validParams, x[0]) {
		return "", "", fmt.Errorf("invalid search parameter")
	}

	return x[0], x[1], nil
}

func URL(link string) string {
	return fmt.Sprintf("http://%s", strings.TrimLeft(link, "/"))
}
