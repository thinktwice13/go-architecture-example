package db

import "fmt"

type Conn struct {
	store map[string][]byte
}

func Connect() *Conn {
	return &Conn{store: make(map[string][]byte)}
}

func (c *Conn) Set(key string, value []byte) {
	c.store[key] = value
	fmt.Printf("Saved: %s\n", string(value))
}

func (c *Conn) Get(key string) ([]byte, bool) {
	value, exists := c.store[key]
	return value, exists
}
