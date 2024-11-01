package database

import (
	"fmt"

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
func (p *PostGres) Delete(id string) (int64, error) {
	query := `UPDATE bookings SET deleted = true WHERE id = $1`

	result, err := p.Repo.Exec(query, id)
	if err != nil {
		return 0, fmt.Errorf("could not mark booking as deleted: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("could not get rows affected: %w", err)
	}

	return rowsAffected, nil
}

// GetLaunchPad gets a launchpad by id.
func (p *PostGres) GetLaunchPad(id string) (*bookings.LaunchPad, error) {
	var result bookings.LaunchPad

	err := p.Repo.QueryRow(`SELECT id, full_name, spacex_launchpad_id, created_at, updated_at FROM launchpads WHERE id = $1`, id).
		Scan(
			&result.Id,
			&result.FullName,
			&result.SpaceXLaunchPadId,
			&result.CreatedAt,
			&result.UpdatedAt,
		)
	if err != nil {
		return nil, fmt.Errorf("error scanning launchpad: %w", err)
	}

	return &result, nil
}

// IsLaunchScheduleValid returns true if there is a launch from the requesed launch pad, day of the week, and destination.
func (p *PostGres) IsLaunchScheduleValid(launchPadId, dayOfWeek, destinationId string) (bool, error) {
	var count int

	err := p.Repo.QueryRow(`SELECT count(*)	FROM launchpad_schedule WHERE launchpad_id = $1 AND destination_id = $2 AND day_of_week = $3`,
		launchPadId, destinationId, dayOfWeek).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error scanning launchpad_schedule: %w", err)
	}

	return count == 1, nil
}
