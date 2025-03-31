package output

import (
	"fmt"
	"hex/core/domain"
	outport "hex/port/output"
)

var _ outport.EventPublisher = (*EventBus)(nil)

type EventBus struct{}

func (e *EventBus) Publish(event domain.DocumentEvent) {
	fmt.Println("Publishing event:", event)
}
