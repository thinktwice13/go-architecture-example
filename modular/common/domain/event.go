package domain

import "time"

type EventType string

// DomainEvent is the base interface for all domain events
type DomainEvent interface {
	EventName() string
	OccurredAt() time.Time
}

type BaseEvent struct {
	Type      EventType
	Timestamp time.Time
}
