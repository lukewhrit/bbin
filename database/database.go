package database

import (
	badger "github.com/dgraph-io/badger/v2"
)

// DB represents an open connection to the Badger database
var DB *badger.DB

// Open creates a new connection to the database
func Open() error {
	var err error

	DB, err = badger.Open(badger.DefaultOptions("/tmp/badger"))

	return err
}

// Close ends the connection to the database
func Close() error {
	return DB.Close()
}
