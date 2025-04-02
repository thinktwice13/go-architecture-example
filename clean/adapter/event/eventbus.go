package event

import (
	"clean/adapter/logging"
	"fmt"

	"clean/domain/event"
)

// Bus implements the event publisher interface defined in the domain
type Bus struct {
	logger logging.Logger
}

func NewEventBus(logger logging.Logger) *Bus {
	return &Bus{
		logger: logger,
	}
}

func (eb *Bus) Publish(event event.DocumentEvent) error {
	fmt.Printf("Publishing event: %s\n", event)
	return nil
}
