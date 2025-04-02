package event

import (
	"fmt"

	"clean/domain/event"
)

var _ event.Publisher = (*Bus)(nil)

type Bus struct{}

func (eb *Bus) Publish(event event.DocumentEvent) error {
	fmt.Printf("Publishing event: %s\n", event)
	return nil
}
