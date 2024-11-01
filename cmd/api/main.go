package main

import (
	"log"
	"net"
	nethttp "net/http"
	"time"

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

	client := &nethttp.Client{
		Timeout: time.Duration(cfg.HTTPTimeout) * time.Second,
		Transport: &nethttp.Transport{
			MaxIdleConns:      cfg.MaxIdleConns,
			MaxConnsPerHost:   cfg.MaxConnsPerHost,
			IdleConnTimeout:   time.Duration(cfg.IdleConnTimeoutSecs) * time.Second,
			DisableKeepAlives: cfg.DisableKeepAlives,
			DialContext: (&net.Dialer{
				Timeout:   time.Duration(cfg.DialerTimeoutSecs) * time.Second,
				KeepAlive: time.Duration(cfg.DialerKeepAliveSecs) * time.Second,
			}).DialContext,
			TLSHandshakeTimeout: time.Duration(cfg.TLSHandshakeTimeoutSecs) * time.Second,
		},
	}

	handlers := api.NewBookingHandlers(repo, client)

	svr := http.New(":8080", handlers)

	log.Printf("API running at http://localhost%s/api/v1\n", ":8080")
	log.Printf("Swagger server running at http://localhost%s\n", ":8081")
	if err := svr.HTTPServer.ListenAndServe(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
