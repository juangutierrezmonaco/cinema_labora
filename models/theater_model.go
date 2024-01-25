package models

import "time"

type Theater struct {
	ID         *int       `json:"id,omitempty"`
	Name       *string    `json:"name,omitempty"`
	Capacity   *int       `json:"capacity,omitempty"`
	LastRow    *string    `json:"last_row,omitempty"`
	LastColumn *int       `json:"last_column,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
}
