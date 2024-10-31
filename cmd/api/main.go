package main

import (
	"log"

	"github.com/petherin/spacetickets/internal/infrastructure/http"
	"github.com/petherin/spacetickets/internal/interfaces/api"
)

func main() {
	handlers := api.NewBookingHandlers(nil)

	svr := http.New(":8080", handlers)

	log.Printf("API running at http://localhost%s/api/v1\n", ":8080")
	log.Printf("Swagger server running at http://localhost%s\n", ":8081")
	if err := svr.HTTPServer.ListenAndServe(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
