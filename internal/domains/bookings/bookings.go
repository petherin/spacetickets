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

type Booker interface {
	GetAll() ([]Booking, error)
	Create(booking Booking) (*Booking, error)
	Delete(bookingId string) (int64, error)
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
