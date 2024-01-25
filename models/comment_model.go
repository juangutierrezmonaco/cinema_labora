package models

type Comment struct {
	ID        *int    `json:"id,omitempty"`
	UserID    *int    `json:"user_id,omitempty"`
	MovieID   *int    `json:"movie_id,omitempty"`
	Content   *string `json:"content,omitempty"`
	CreatedAt *int64  `json:"created_at,omitempty"`
}
