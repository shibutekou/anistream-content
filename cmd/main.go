package main

import (
	"github.com/vgekko/ani-go/config"
	"github.com/vgekko/ani-go/internal/app"
	"log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal("config error: ", err.Error())
	}

	app.Run(cfg)
}
