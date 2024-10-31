package api

import (
	"net/http"

	"github.com/petherin/spacetickets/internal/domains/bookings"
)

type BookingHandlers struct {
	Booker bookings.Booker
}

func NewBookingHandlers(booker bookings.Booker) BookingHandlers {
	return BookingHandlers{Booker: booker}
}

func (b *BookingHandlers) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get"))
}

func (b *BookingHandlers) Post(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Post"))
}

func (b *BookingHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete"))
}
