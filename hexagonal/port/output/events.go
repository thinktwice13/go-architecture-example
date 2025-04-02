package output

import (
	"hexagonal/core/domain"
)

// EventPublisher defines the output port for publishing events
// This is another hexagon boundary that core services use to communicate outward
type EventPublisher interface {
	Publish(domain.DocumentEvent) error
}
