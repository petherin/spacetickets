package main

import (
	"log"

	"github.com/petherin/spacetickets/internal/infrastructure/config"
	"github.com/petherin/spacetickets/internal/infrastructure/database"
	"github.com/petherin/spacetickets/internal/infrastructure/http"
	"github.com/petherin/spacetickets/internal/interfaces/api"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		log.Fatalf("failed to get config: %s\n", err)
	}

	repo, err := database.New(cfg)
	if err != nil {
		log.Fatalf("failed to create booking repo client: %s\n", err)
	} 
	defer repo.Close()

	handlers := api.NewBookingHandlers(repo)

	svr := http.New(":8080", handlers)

	log.Printf("API running at http://localhost%s/api/v1\n", ":8080")
	log.Printf("Swagger server running at http://localhost%s\n", ":8081")
	if err := svr.HTTPServer.ListenAndServe(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
