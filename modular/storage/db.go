package storage

import "fmt"

type DB struct {
	store map[string]any
}

func ConnectDB() *DB {
	// Simulate a database connection
	return &DB{store: make(map[string]any)}
}

func (db *DB) Set(key string, value any) error {
	// Simulate inserting a record into the database
	db.store[key] = value
	return nil
}

func (db *DB) Get(key string) (any, error) {
	// Simulate retrieving a record from the database
	if value, exists := db.store[key]; exists {
		return value, nil
	}
	return nil, fmt.Errorf("record not found")
}
