package db

type DB interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
}

type InMemoryDB struct {
	data map[string]interface{}
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		data: make(map[string]interface{}),
	}
}

func (db *InMemoryDB) Set(key string, value interface{}) error {
	db.data[key] = value
	return nil
}

func (db *InMemoryDB) Get(key string) (interface{}, error) {
	value, exists := db.data[key]
	if !exists {
		return nil, nil // or return an error if preferred
	}
	return value, nil
}
