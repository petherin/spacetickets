package bookings

import "time"

type Booking struct {
	Id string
	Customer
	LaunchPadId   string
	DestinationId string
	LaunchDate    time.Time
	Deleted       bool
}

type Customer struct {
	FirstName string
	LastName  string
	Gender    string
	Birthday  time.Time
}

type Booker interface {
	GetAll() ([]Booking, error)
	Create(booking Booking) (*Booking, error)
	Delete(bookingId string) error
}
