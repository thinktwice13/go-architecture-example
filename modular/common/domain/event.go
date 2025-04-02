package domain

import "time"

type EventType string

var (
	EventDocumentUploaded EventType = "document.uploaded"
)

// Event is the base interface for all domain events
type Event interface {
	EventName() string
	OccurredAt() time.Time
}

type BaseEvent struct {
	Type      EventType `json:"type"`
	Timestamp time.Time `json:"timestamp"`
}
