package models

import "time"

type User struct {
	ID           *int       `json:"id,omitempty"`
	FirstName    *string    `json:"first_name,omitempty"`
	LastName     *string    `json:"last_name,omitempty"`
	Email        *string    `json:"email,omitempty"`
	Password     *string    `json:"password,omitempty"`
	PasswordHash string     `json:"-"`
	Gender       *string    `json:"gender,omitempty"`
	PictureURL   *string    `json:"picture_url,omitempty"`
	TicketIDs    []int      `json:"ticket_ids,omitempty"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
