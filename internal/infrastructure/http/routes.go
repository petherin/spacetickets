package http

import (
	"net/http"

	"github.com/petherin/spacetickets/internal/interfaces/api"
)

// NewMux sets up routes for the API.
func (s *Server) NewMux(handlers api.BookingHandlers) *http.ServeMux {
	const baseURL = "/api/v1"

	mux := http.NewServeMux()
	mux.HandleFunc("GET "+baseURL+"/bookings", handlers.Get)
	mux.HandleFunc("POST "+baseURL+"/booking", handlers.Post)
	mux.HandleFunc("DELETE "+baseURL+"/booking/{id}", handlers.Delete)

	return mux
}
