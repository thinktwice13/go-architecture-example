package logging

import (
	"clean/usecase/iface"
	"fmt"
)

var _ iface.Logger = (*Logger)(nil)

type Logger struct{}

func (l *Logger) Log(message string) {
	fmt.Printf("INFO: %s\n", message)
}
