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
	service := domain.NewService(token, http.Client{})
	handler := delivery.NewHandler(service, log)

	app := handler.InitRoutes()

	if err := app.Listen(":8800"); err != nil {
		log.Fatal(err.Error())
	}
}
