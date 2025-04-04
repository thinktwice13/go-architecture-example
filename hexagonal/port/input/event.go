package input

import "hexagonal/core/domain"

type Handler interface {
	Handle(domain.DocumentEvent) error
}

type HandlerFunc func(eventType string, eventData interface{}) error

type EventSubscriber interface {
	Subscribe(eventType string, h Handler) error
}
