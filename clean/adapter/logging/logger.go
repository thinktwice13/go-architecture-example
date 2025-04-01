package logging

import "fmt"

// Logger defines logging interface
type Logger interface {
	Log(message string)
}

type DefaultLogger struct{}

func (DefaultLogger) Log(message string) {
	fmt.Println(message)
}
