package http

import (
	"net/http"
	"time"

	"github.com/petherin/spacetickets/internal/interfaces/api"
)

// Server encapsulates an HTTP server.
type Server struct {
	HTTPServer *http.Server
}

// New creates a new server, setting its address and handlers to those passed in.
func New(addr string, handlers api.BookingHandlers) Server {
	server := Server{}

	mux := server.NewMux(handlers)
	mw := server.RecoverPanic(server.LogRequest(server.CORS(mux)))
	svr := http.Server{
		Addr:         addr,
		Handler:      mw,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	server.HTTPServer = &svr

	return server
}
