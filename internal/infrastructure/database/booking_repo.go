package database

import (
	"fmt"
	"log"

	"github.com/petherin/spacetickets/internal/domains/bookings"
)

// Get returns all bookings that aren't marked as deleted.
func (p *PostGres) GetAll() ([]bookings.Booking, error) {
	rows, err := p.Repo.Query(`SELECT id, first_name, last_name, gender, birthday, launchpad_id, destination_id, launch_date, created_at, updated_at FROM bookings WHERE deleted = false`)
	if err != nil {
		return nil, fmt.Errorf("error querying bookings: %w", err)
	}
	defer rows.Close()

	results := []bookings.Booking{}

	for rows.Next() {
		var result bookings.Booking
		if err := rows.Scan(
			&result.Id,
			&result.FirstName,
			&result.LastName,
			&result.Gender,
			&result.Birthday,
			&result.LaunchPadId,
			&result.DestinationId,
			&result.LaunchDate,
			&result.CreatedAt,
			&result.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("error scanning bookings: %w", err)
		}
		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over bookings rows: %w", err)
	}

	return results, nil
}

// Get retrieves the requested booking.
func (p *PostGres) Get(id string) (*bookings.Booking, error) {
	var result bookings.Booking

	err := p.Repo.QueryRow(`SELECT id, first_name, last_name, gender, birthday, launchpad_id, destination_id, launch_date, created_at, updated_at FROM bookings WHERE id = $1`, id).
		Scan(
			&result.Id,
			&result.FirstName,
			&result.LastName,
			&result.Gender,
			&result.Birthday,
			&result.LaunchPadId,
			&result.DestinationId,
			&result.LaunchDate,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
	if err != nil {
		return nil, fmt.Errorf("error scanning booking: %w", err)
	}

	return &result, nil
}

// Create adds a new booking.
func (p *PostGres) Create(booking bookings.Booking) (*bookings.Booking, error) {
	var insertedID string

	err := p.Repo.QueryRow(`INSERT INTO bookings (first_name, last_name, gender, birthday, launchpad_id, destination_id, launch_date, created_at, updated_at)
	 VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW()) RETURNING id`,
		booking.FirstName, booking.LastName, booking.Gender, booking.Birthday, booking.LaunchPadId, booking.DestinationId, booking.LaunchDate).Scan(&insertedID)
	if err != nil {
		return nil, fmt.Errorf("error creating booking: %w", err)
	}

	return p.Get(insertedID)
}

// Delete marks a booking as deleted.
func (p *PostGres) Delete(id string) error {
	query := `UPDATE bookings SET deleted = true WHERE id = $1`

	result, err := p.Repo.Exec(query, id)
	if err != nil {
		return fmt.Errorf("could not mark booking as deleted: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("could not get rows affected: %w", err)
	}

	log.Printf("Number of rows updated: %d\n", rowsAffected)

	return nil
}