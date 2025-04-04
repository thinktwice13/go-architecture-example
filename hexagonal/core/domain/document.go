package domain

import (
	"time"
)

type Document struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

// Validate ensures the document meets domain rules
// Note: Pure domain logic with no dependencies
func (d *Document) Validate() error {
	if d.Name == "" {
		return ErrInvalidDocument("document name cannot be empty")
	}

	return nil
}
