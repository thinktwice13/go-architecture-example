package model

import "time"

type Document struct {
	ID        string
	Name      string
	Status    string
	CreatedAt time.Time
}
