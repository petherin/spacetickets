package bookings

import (
	"encoding/json"
	"fmt"
	"time"
)

type Booking struct {
	Id string `json:"id"`
	Customer
	LaunchPadId   string    `json:"launch_pad_id"`
	DestinationId string    `json:"destination_id"`
	LaunchDate    time.Time `json:"launch_date"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Customer struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Gender    string    `json:"gender"`
	Birthday  time.Time `json:"birthday"`
}

type LaunchPad struct {
	Id                string    `json:"id"`
	FullName          string    `json:"full_name"`
	SpaceXLaunchPadId string    `json:"spacex_launchpad_id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type SpaceXLaunches struct {
	TotalDocs int `json:"totalDocs"`
}

type SpaceXLaunchesRequest struct {
	Query           Query   `json:"query"`
	Options         Options `json:"options"`
	ResolveBodyOnly bool    `json:"resolveBodyOnly"`
	ResponseType    string  `json:"responseType"`
}

type Query struct {
	LaunchPad string  `json:"launchpad"`
	DateUtc   DateUtc `json:"date_utc"`
}
type DateUtc struct {
	Gte time.Time `json:"$gte"`
	Lt  time.Time `json:"$lt"`
}

type Options struct {
	Pagination bool `json:"pagination"`
	Limit      int  `json:"limit"`
}

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
