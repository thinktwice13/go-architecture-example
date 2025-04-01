package event

import (
	"fmt"
	"hex/core/domain"
	"hex/port/output"
)

// Bus implements EventPublisher
type Bus struct{}

var _ output.EventPublisher = (*Bus)(nil)

// Publish publishes an event to the event bus
func (*Bus) Publish(_ domain.DocumentEvent) error {
	fmt.Println("Publishing event to the bus")
	return nil
}
