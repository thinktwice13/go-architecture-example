package event

import (
	"fmt"
	"vertical/shared/domain"
)

type Publisher interface {
	Publish(domain.DocumentEvent) error
}

type Bus struct{}

func (b *Bus) Publish(event domain.DocumentEvent) error {
	fmt.Println("Publishing event:", event)
	return nil
}
