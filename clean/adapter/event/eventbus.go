package event

import (
	"fmt"

	"clean/domain/event"
)

type Bus struct{}

func (eb *Bus) Publish(event event.DocumentEvent) error {
	fmt.Printf("Publishing event: %s\n", event)
	return nil
}
