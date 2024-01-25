package models

type Screening struct {
	ID             *int     `json:"id,omitempty"`
	Name           *string  `json:"name,omitempty"`
	MovieID        *int     `json:"movie_id,omitempty"`
	TheaterID      *int     `json:"theater_id,omitempty"`
	AvailableSeats *int     `json:"available_seats,omitempty"`
	TakenSeats     []string `json:"taken_seats,omitempty"`
	Showtime       *int64   `json:"showtime,omitempty"`
	Price          *float64 `json:"price,omitempty"`
	Language       *string  `json:"language,omitempty"`
	ViewsCount     *int     `json:"views_count,omitempty"`
	CreatedAt      *int64   `json:"created_at,omitempty"`
	UpdatedAt      *int64   `json:"updated_at,omitempty"`
}
