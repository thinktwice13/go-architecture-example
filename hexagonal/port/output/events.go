package output

import "hex/core/domain"

type EventPublisher interface {
	Publish(event domain.DocumentEvent)
}
