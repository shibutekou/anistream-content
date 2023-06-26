package normalize

import (
	"fmt"
	"strings"
)

func Params(l, sep string) (string, string) {
	x := strings.Split(l, sep)
	return x[0], x[1]
}

func Link(link string) string {
	return fmt.Sprintf("http://%s", strings.TrimLeft(link, "/"))
}
