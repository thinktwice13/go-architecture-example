package db

type DB struct{}

func ConnectDB() *DB {
	return &DB{}
}
