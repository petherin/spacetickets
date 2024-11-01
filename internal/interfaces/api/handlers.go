package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/petherin/spacetickets/internal/domains/bookings"
)

// BookingHandlers provides methods and dependencies needed to handle requests to the API.
type BookingHandlers struct {
	Booker            bookings.Booker
	HTTPClient        *http.Client
	SpaceXAPIEndpoint string
}

// NewBookingHandlers returns a new BookingHandlers object, assigning passed dependencies.
func NewBookingHandlers(booker bookings.Booker, client *http.Client, spaceXAPIEndpoint string) BookingHandlers {
	return BookingHandlers{Booker: booker, HTTPClient: client, SpaceXAPIEndpoint: spaceXAPIEndpoint}
}

// Get returns all bookings.
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

// Post validates the requested booking and creates it if so.
func (b *BookingHandlers) Post(w http.ResponseWriter, r *http.Request) {
	var booking bookings.Booking

	err := json.NewDecoder(r.Body).Decode(&booking)
	if err != nil {
		log.Println(err)
		http.Error(w, "an error occurred, see logs", http.StatusInternalServerError)
		return
	}

	launchPad, err := b.Booker.GetLaunchPad(booking.LaunchPadId)
	if err != nil {
		log.Println(err)
		http.Error(w, "an error occurred, see logs", http.StatusInternalServerError)
		return
	}

	spaceXLaunches, err := b.getSpaceXLaunch(launchPad.SpaceXLaunchPadId, booking)
	if err != nil {
		log.Println(err)
		http.Error(w, "an error occurred, see logs", http.StatusInternalServerError)
		return
	}

	if spaceXLaunches.TotalDocs > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"Status": "Flight cancelled, overlaps with SpaceX launch"}`))
		return
	}

	proposedWeekDay := booking.LaunchDate.Weekday().String()
	validLaunch, err := b.Booker.IsLaunchScheduleValid(booking.LaunchPadId, proposedWeekDay, booking.DestinationId)
	if err != nil {
		log.Println(err)
		http.Error(w, "an error occurred, see logs", http.StatusInternalServerError)
		return
	}

	if !validLaunch {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"Status": "Flight cancelled, this launchpad does not fly to the destination on the requested day"}`))
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

// Delete marks the specified booking as deleted.
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

// getSpaceXLaunch contacts the SpaceX API to check if there is a SpaceX launch from the requested launchpad and date.
func (b *BookingHandlers) getSpaceXLaunch(spaceXLaunchId string, booking bookings.Booking) (*bookings.SpaceXLaunches, error) {
	fullURL, err := url.JoinPath(b.SpaceXAPIEndpoint, "/v4/launches/query")
	if err != nil {
		return nil, err
	}

	payload := bookings.SpaceXLaunchesRequest{
		Query: bookings.Query{
			LaunchPad: spaceXLaunchId,
			DateUtc: bookings.DateUtc{
				Gte: booking.LaunchDate,
				Lt:  booking.LaunchDate.Add(24 * time.Hour),
			},
		},
		Options: bookings.Options{
			Pagination: true,
			Limit:      0,
		},
		ResolveBodyOnly: true,
		ResponseType:    "json",
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := b.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var spaceXLaunches bookings.SpaceXLaunches
	err = json.NewDecoder(resp.Body).Decode(&spaceXLaunches)
	if err != nil {
		return nil, err
	}

	return &spaceXLaunches, nil
}
