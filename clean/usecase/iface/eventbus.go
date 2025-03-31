package iface

import "clean/entity"

type EventBus interface {
	Publish(event entity.DocumentEvent) error
}
