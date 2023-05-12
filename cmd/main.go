package main

import (
	"github.com/vgekko/ani-go/internal/delivery"
	"github.com/vgekko/ani-go/internal/domain"
	"github.com/vgekko/ani-go/pkg/logger"
	"net/http"
	"os"
)

func main() {
	token := os.Getenv("KODIK_TOKEN")

	log := logger.GetLogger()
	link := domain.NewLink(token, http.Client{})
	info := domain.NewInfo(token, http.Client{})
	handler := delivery.NewHandler(link, info, log)

	app := handler.InitRoutes()

	if err := app.Listen(":8800"); err != nil {
		log.Fatal(err.Error())
	}
}
