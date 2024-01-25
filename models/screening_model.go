package models

import "time"

type Screening struct {
	ID             *int       `json:"id,omitempty"`
	Name           *string    `json:"name,omitempty"`
	MovieID        *int       `json:"movie_id,omitempty"`
	TheaterID      *int       `json:"theater_id,omitempty"`
	AvailableSeats *int       `json:"available_seats,omitempty"`
	TakenSeats     []string   `json:"taken_seats,omitempty"`
	Showtime       *time.Time `json:"showtime,omitempty"`
	Price          *float64   `json:"price,omitempty"`
	Language       *string    `json:"language,omitempty"`
	ViewsCount     *int       `json:"views_count,omitempty"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time  `json:"updated_at,omitempty"`
}
