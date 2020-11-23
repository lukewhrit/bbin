package database

import (
	badger "github.com/dgraph-io/badger/v2"
)

// Set sets a value in the database
func Set(key, value []byte) error {
	return DB.Update(func(txn *badger.Txn) error {
		return txn.Set(key, value)
	})
}

// Get retrieves a value in the database
func Get(key []byte) ([]byte, error) {
	var value []byte

	err := DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)

		if err != nil {
			return err
		}

		value, err = item.ValueCopy(nil)

		return err
	})

	return value, err
}

// Delete removes an entry from the database
func Delete(key []byte) error {
	return DB.Update(func(txn *badger.Txn) error {
		return txn.Delete(key)
	})
}
