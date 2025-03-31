package infra

type DB struct {
	store map[string]any
}

func ConnectDB() *DB {
	return &DB{
		store: make(map[string]any),
	}
}

func (db *DB) Get(key string) (any, bool) {
	value, ok := db.store[key]
	return value, ok
}

func (db *DB) Set(key string, value any) {
	db.store[key] = value
}
