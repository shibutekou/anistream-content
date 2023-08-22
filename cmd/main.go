package main

import (
	"github.com/vgekko/anistream-content/config"
	"github.com/vgekko/anistream-content/internal/app"
)

func main() {
	cfg := config.Load()

	app.Run(cfg)
}
