package models

type Ticket struct {
	ID          *int    `json:"id,omitempty"`
	PickupID    *string `json:"pickup_id,omitempty"`
	UserID      *int    `json:"user_id,omitempty"`
	ScreeningID *int    `json:"screening_id,omitempty"`
	CreatedAt   *int64  `json:"created_at,omitempty"`
}
