package main

import (
	"github.com/vgekko/ani-go/internal/delivery"
	"github.com/vgekko/ani-go/internal/domain"
	"log"
	"net/http"
	"os"
)

func main() {
	token := os.Getenv("KODIK_TOKEN")

	kodik := domain.NewKodik(http.Client{}, token)
	handler := delivery.NewHandler(kodik)

	app := handler.InitRoutes()

	if err := app.Listen(":8800"); err != nil {
		log.Fatal(err)
	}
}
