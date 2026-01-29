package models

import "time"

type Company struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	Name       string    `json:"name"`
	Slug       string    `json:"slug"`
	Components string    `json:"components"` // JSONB as string
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
