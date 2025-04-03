package model

import "time"

type Document struct {
	ID        int64
	Name      string
	Status    string
	CreatedAt time.Time
}
