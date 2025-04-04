package event

import (
	"clean/domain/event"
	"fmt"
)

var _ event.Publisher = (*Bus)(nil)

type Bus struct{}

func (eb *Bus) Publish(event event.DocumentEvent) error {
	fmt.Printf("Event \"%s\" published\n", event.Type)
	return nil
}
