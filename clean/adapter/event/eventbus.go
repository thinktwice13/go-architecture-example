package event

import (
	"fmt"

	"clean/domain/event"
)

// Publisher is an example of a convenience interface defined as a helper for usecases
type Publisher interface {
	Publish(event event.DocumentEvent) error
}

type Bus struct{}

func (eb *Bus) Publish(event event.DocumentEvent) error {
	fmt.Printf("Publishing event: %s\n", event)
	return nil
}
