package db

// type DB interface {
// 	Set(key string, value interface{}) error
// 	Get(key string) (interface{}, error)
// }

type DB struct {
	data map[string]interface{}
}

func NewDB() *DB {
	return &DB{
		data: make(map[string]interface{}),
	}
}

func (db *DB) Set(key string, value interface{}) error {
	db.data[key] = value
	return nil
}

func (db *DB) Get(key string) (interface{}, error) {
	value, exists := db.data[key]
	if !exists {
		return nil, nil // or return an error if preferred
	}
	return value, nil
}
