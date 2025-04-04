package eventbus

import (
	"fmt"
	"hexagonal/core/domain"
	"hexagonal/port/input"
	"hexagonal/port/output"
)

// Bus implements two ports (input and output)
type Bus struct{}

var (
	_ output.EventPublisher = (*Bus)(nil)
	_ input.EventSubscriber = (*Bus)(nil)
)

func (*Bus) Publish(domain.DocumentEvent) error {
	fmt.Println("Publishing event to the bus")
	return nil
}

func (*Bus) Subscribe(eventType string, _ input.Handler) error {
	fmt.Printf("Subscribed to events: %s", eventType)
	return nil
}
