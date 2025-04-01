package document

import (
	"clean/domain/event"
)

type EventBus interface {
	Publish(event event.DocumentEvent) error
}
