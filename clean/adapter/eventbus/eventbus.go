package eventbus

import (
	"clean/adapter/logging"
	"fmt"

	"clean/domain/event"
)

// EventBus implements the event publisher interface defined in the domain
type EventBus struct {
	logger logging.Logger
}

func NewEventBus(logger logging.Logger) *EventBus {
	return &EventBus{
		logger: logger,
	}
}

func (eb *EventBus) Publish(event event.DocumentEvent) error {
	fmt.Printf("Publishing event: %s\n", event)
	return nil
}
