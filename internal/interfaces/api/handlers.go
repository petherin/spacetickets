package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/petherin/spacetickets/internal/domains/bookings"
)

type BookingHandlers struct {
	Booker     bookings.Booker
	HTTPClient *http.Client
}

func NewBookingHandlers(booker bookings.Booker, client *http.Client) BookingHandlers {
	return BookingHandlers{Booker: booker, HTTPClient: client}
}

func (b *BookingHandlers) Get(w http.ResponseWriter, r *http.Request) {
	bookings, err := b.Booker.GetAll()
	if err != nil {
		log.Println(err)
		http.Error(w, "an error occurred, see logs", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(bookings)
	if err != nil {
		log.Println(err)
		http.Error(w, "an error occurred, see logs", http.StatusInternalServerError)
	}

}

func (b *BookingHandlers) Post(w http.ResponseWriter, r *http.Request) {
	var booking bookings.Booking

	err := json.NewDecoder(r.Body).Decode(&booking)
	if err != nil {
		log.Println(err)
		http.Error(w, "an error occurred, see logs", http.StatusInternalServerError)
		return
	}

	newBooking, err := b.Booker.Create(booking)
	if err != nil {
		log.Println(err)
		http.Error(w, "an error occurred, see logs", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(newBooking)
	if err != nil {
		log.Println(err)
		http.Error(w, "an error occurred, see logs", http.StatusInternalServerError)
		return
	}
}

func (b *BookingHandlers) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if len(id) == 0 {
		http.NotFound(w, r)
		return
	}

	rowsAffected, err := b.Booker.Delete(id)
	if err != nil {
		log.Println(err)
		http.Error(w, "an error occurred, see logs", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	log.Printf("Number of rows updated: %d\n", rowsAffected)

	if rowsAffected == 0 {
		w.Write([]byte(`{"Status": "ID not recognised"}`))
		return
	}

	w.Write([]byte(`{"Status": "Record deleted"}`))
}
