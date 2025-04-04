package event

import (
	"fmt"
	"vertical/shared/domain"
)

type Bus struct{}

func (eb *Bus) Publish(event domain.DocumentEvent) error {
	fmt.Printf("Event \"%s\" published\n", event.Type)
	return nil
}
