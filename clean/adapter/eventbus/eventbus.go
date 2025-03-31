package eventbus

import (
	"clean/entity"
	"clean/usecase/iface"
	"fmt"
)

var _ iface.EventBus = (*Bus)(nil)

type Bus struct{}

func (eb *Bus) Publish(event entity.DocumentEvent) error {
	fmt.Println("Publishing event:", event)
	return nil
}
