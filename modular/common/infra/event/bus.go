package event

import (
	"fmt"
	"modular/common/domain"
)

type Bus struct{}

func (*Bus) Publish(domain.Event) error {
	fmt.Println("Publishing event...")
	return nil
}
