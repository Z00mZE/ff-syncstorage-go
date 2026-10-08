package types

import "time"

type Collection struct {
	ID        uint64    `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}
