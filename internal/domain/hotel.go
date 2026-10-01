package domain

import "time"

type Hotel struct {
	ID        int       `json:"id"`
	Name      string    `json:"name,omitempty"`
	City      string    `json:"city,omitempty"`
	Address   string    `json:"address,omitempty"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}
