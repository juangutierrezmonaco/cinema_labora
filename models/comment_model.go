package models

import "time"

type Comment struct {
	ID        *int       `json:"id,omitempty"`
	UserID    *int       `json:"user_id,omitempty"`
	MovieID   *int       `json:"movie_id,omitempty"`
	Content   *string    `json:"content,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}
