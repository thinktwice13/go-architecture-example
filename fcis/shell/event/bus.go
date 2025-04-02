package event

import (
	"fcis/core/document"
	"fmt"
)

type Publisher interface {
	Publish(event document.Event) error
}

type Bus struct{}

func (b *Bus) Publish(event document.Event) error {
	fmt.Println("Publishing event:", event)
	return nil
}
