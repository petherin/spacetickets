package bookings

import (
	"encoding/json"
	"fmt"
	"time"
)

// Booking represents a flight booking.
type Booking struct {
	Id string `json:"id"`
	Customer
	LaunchPadId   string    `json:"launch_pad_id"`
	DestinationId string    `json:"destination_id"`
	LaunchDate    time.Time `json:"launch_date"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// Customer represents a customer wishing to book a flight.
type Customer struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Gender    string    `json:"gender"`
	Birthday  time.Time `json:"birthday"`
}

// LaunchPad represents a launch pad name and id, and maps to the corresponding launch id that SpaceX use.
type LaunchPad struct {
	Id                string    `json:"id"`
	FullName          string    `json:"full_name"`
	SpaceXLaunchPadId string    `json:"spacex_launchpad_id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// SpaceXLaunches lists how many launches are found.
type SpaceXLaunches struct {
	TotalDocs int `json:"totalDocs"`
}

// SpaceXLaunchesRequest allows us to make requests to the Space X API.
type SpaceXLaunchesRequest struct {
	Query           Query   `json:"query"`
	Options         Options `json:"options"`
	ResolveBodyOnly bool    `json:"resolveBodyOnly"`
	ResponseType    string  `json:"responseType"`
}

// Query lists the fields we want to query on the SpaceX API.
type Query struct {
	LaunchPad string  `json:"launchpad"`
	DateUtc   DateUtc `json:"date_utc"`
}

// DateUtc lets us denote a date range when querying SpaceX API.
type DateUtc struct {
	Gte time.Time `json:"$gte"`
	Lt  time.Time `json:"$lt"`
}

// Options are pagination options when querying SpaceX API.
type Options struct {
	Pagination bool `json:"pagination"`
	Limit      int  `json:"limit"`
}

// Booker defines the methods an object needs to implement to list, create, delete and validate bookings.
type Booker interface {
	GetAll() ([]Booking, error)
	Create(booking Booking) (*Booking, error)
	Delete(bookingId string) (int64, error)
	GetLaunchPad(id string) (*LaunchPad, error)
	IsLaunchScheduleValid(launchPadId, dayOfWeek, destinationId string) (bool, error)
}

// UnmarshalJSON unmarshals booking JSON so that dates have the proper time.Time format.
func (r *Booking) UnmarshalJSON(data []byte) error {
	type Alias Booking
	aux := &struct {
		Birthday   string `json:"birthday"`
		LaunchDate string `json:"launch_date"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	dateOnly, err := time.Parse(time.DateOnly, aux.Birthday)
	if err != nil {
		return fmt.Errorf("invalid date format. Use YYYY-MM-DD: %w", err)
	}

	r.Birthday = dateOnly

	dateOnly, err = time.Parse(time.DateOnly, aux.LaunchDate)
	if err != nil {
		return fmt.Errorf("invalid date format. Use YYYY-MM-DD: %w", err)
	}

	r.LaunchDate = dateOnly

	return nil
}
