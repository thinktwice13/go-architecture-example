package entity

import "time"

type Document struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

func (d *Document) ValidateForUpload() error {
	if d.Name == "" {
		return ErrInvalidDocument{Field: "name", Message: "document name cannot be empty"}
	}

	return nil
}
