package http

import (
	"net/http"
	"time"

	"github.com/petherin/spacetickets/internal/interfaces/api"
)

type Server struct {
	HTTPServer *http.Server
}

func New(addr string, handlers api.BookingHandlers) Server {
	server := Server{}

	mux := server.NewMux(handlers)
	svr := http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	server.HTTPServer = &svr

	return server
}
