package domain

import "time"

type Hotel struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	City      string    `json:"city"`
	Address   string    `json:"address"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}
