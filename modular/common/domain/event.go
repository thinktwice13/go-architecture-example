package domain

import "time"

type EventType string

// Event is the base interface for all domain events
type Event struct {
	EventType EventType
	Timestamp time.Time
	Data      []byte
}
