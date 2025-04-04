package event

import (
	"fmt"
	"modular/common/domain"
)

type Bus struct{}

func (*Bus) Publish(event domain.Event) error {
	fmt.Printf("Event \"%s\" published\n", event.EventType)
	return nil
}
