package event

import (
	"fmt"
	"hexagonal/core/domain"
	"hexagonal/port/output"
)

// Bus implements EventPublisher
type Bus struct{}

var _ output.EventPublisher = (*Bus)(nil)

// Publish publishes an event to the event bus
func (*Bus) Publish(domain.DocumentEvent) error {
	fmt.Println("Publishing event to the bus")
	return nil
}
