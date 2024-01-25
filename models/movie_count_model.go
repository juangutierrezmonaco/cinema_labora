package models

type MovieCount struct {
	MovieID    *int   `json:"movie_id"`
	ViewsCount *int   `json:"views_count,omitempty"`
	UpdatedAt  *int64 `json:"updated_at"`
}
